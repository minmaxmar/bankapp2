package kafka

import (
	"bankapp2/app/models"
	"bankapp2/helper/config"
	"context"
	"strconv"
	"time"

	"log/slog"

	"github.com/go-stack/stack"
	"gorm.io/gorm"

	"github.com/IBM/sarama"
)

type kafkaImpl struct {
	logger           *slog.Logger
	producer         sarama.SyncProducer
	consumer         sarama.Consumer
	cardRepo         cardRepo
	bootstrapServers string
	topic            string
}

type cardRepo interface {
	GetConn() *gorm.DB
	BeginTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB)
	RollbackTransaction(tx *gorm.DB)
	DeleteCardID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (int64, error)
	GetExpiredCards(connWithOrNoTx *gorm.DB, ctx context.Context) ([]*models.Card, error)
}

type Kafka interface {
	// ProduceDeleteCard(ctx context.Context, id int) error
	ProduceDeleteExpiredCards(ctx context.Context) error
	ProcessConsumer(ctx context.Context) chan error
	ConsumeCardDelete(ctx context.Context, msg *sarama.ConsumerMessage) (int64, error)
	ScheduleProducer(ctx context.Context) chan error
	StopKafka()
}

func NewConn(cardRepo cardRepo, config config.Config, logger *slog.Logger) (Kafka, error) {
	configSarama := sarama.NewConfig()
	configSarama.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{config.Kafka.BootstrapServers}, configSarama)
	if err != nil {
		logger.Error("Unable to create Kafka producer",
			"error", err,
			"stacktrace", stack.Trace().String())
		return &kafkaImpl{}, err
	}
	// TODO: to boostrap
	// producer.Close()

	consumer, err := sarama.NewConsumer([]string{config.Kafka.BootstrapServers}, nil)
	if err != nil {
		logger.Error("Unable to create Kafka consumer",
			"error", err,
			"stacktrace", stack.Trace().String())
		return &kafkaImpl{}, err
	}
	// TODO: to boostrap
	// defer consumer.Close()
	logger.Info("Consumer started")

	logger.Info("New Kafka connection opened")
	return &kafkaImpl{
		logger:           logger,
		producer:         producer,
		consumer:         consumer,
		cardRepo:         cardRepo,
		bootstrapServers: config.Kafka.BootstrapServers,
		topic:            config.Kafka.ExpiredCardsTopic,
	}, nil
}

func (k *kafkaImpl) StopKafka() {
	k.consumer.Close()
	k.producer.Close()
}

func (k *kafkaImpl) ScheduleProducer(ctx context.Context) chan error {
	ticker := time.NewTicker(1 * time.Minute)
	errChan := make(chan error)
	go func() {
		defer ticker.Stop()
		defer close(errChan)
		for {
			select {
			case <-ticker.C:
				k.logger.Info("Next tick")
				if err := k.ProduceDeleteExpiredCards(ctx); err != nil {
					errChan <- err
					return
				}
			case <-ctx.Done():
				k.logger.Info("Shutting down ProduceDeleteExpiredCards goroutine.")
				return
			}
		}
	}()

	return errChan
}

func (k *kafkaImpl) ProduceDeleteExpiredCards(ctx context.Context) error {
	cards, err := k.cardRepo.GetExpiredCards(k.cardRepo.GetConn(), ctx)
	k.logger.Info("Got exp. cards", "cards", cards)
	if err != nil {
		k.logger.Error("Failed while getting expired cards",
			"error", err,
			"stacktrace", stack.Trace().String())
		return err
	}
	for _, card := range cards {
		message := strconv.Itoa(int(card.ID)) // serialize message
		_, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
			Topic: k.topic,
			Value: sarama.StringEncoder(message),
		})
		k.logger.Info("Sent mes to kafka card ID", "cardID", message)
		if err != nil {
			k.logger.Error("Failed to produce message", "error", err)
			return err
		}
	}
	return nil
}

// TODO: !!! all consumers created get no closed after NewConsumer exists!!!!
func (k *kafkaImpl) ProcessConsumer(ctx context.Context) chan error {
	errChan := make(chan error)
	partitions, err := k.consumer.Partitions(k.topic)
	if err != nil {
		k.logger.Error("Failed to fetch partitions",
			"topic", k.topic,
			"error", err,
			"stacktrace", stack.Trace().String())
		errChan <- err
		return errChan
	}

	for _, partition := range partitions {
		partitionConsumer, err := k.consumer.ConsumePartition(k.topic, partition, sarama.OffsetNewest)

		if err != nil {
			k.logger.Error("Failed to start partition consumer",
				"partition", partition,
				"error", err,
				"stacktrace", stack.Trace().String())
			continue
		}

		go func(pc sarama.PartitionConsumer) {
			defer partitionConsumer.Close()
			for msg := range partitionConsumer.Messages() {
				if deletedID, err := k.ConsumeCardDelete(ctx, msg); err != nil {
					k.logger.Error("Error while trying to delete", "cardID", msg.Value, "error", err)
					errChan <- err
					return
				} else {
					k.logger.Info("Successfully deleted card ID", "deletedID", deletedID)
				}
			}
		}(partitionConsumer)
	}
	return errChan
}

func (k *kafkaImpl) ConsumeCardDelete(ctx context.Context, msg *sarama.ConsumerMessage) (int64, error) {
	msgString := string(msg.Value)
	cardID, err := strconv.ParseInt(msgString, 10, 64)
	if err != nil {
		k.logger.Error("Error parsing cardID from Kafka message", "error", err)
		return 0, err
	}
	k.logger.Info("Received message to delete card ID", "card ID", msgString)
	tx := k.cardRepo.BeginTransaction()
	deletedID, err := k.cardRepo.DeleteCardID(tx, ctx, cardID)
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return deletedID, err
}
