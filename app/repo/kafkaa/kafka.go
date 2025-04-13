package kafka

import (
	"bankapp2/app/models"
	"bankapp2/helper/config"
	"context"
	"strconv"
	"time"

	"log/slog"

	"github.com/go-stack/stack"
	"github.com/robfig/cron"
	"gorm.io/gorm"

	"github.com/IBM/sarama"
)

type kafkaProducer struct {
	logger           *slog.Logger
	producer         sarama.SyncProducer
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
	NewConsumer(ctx context.Context) error
	ConsumeCardDelete(ctx context.Context, msg *sarama.ConsumerMessage) (int64, error)
	ScheduleProducer(ctx context.Context) error
	ScheduleConsumer(ctx context.Context) error
}

func NewConn(cardRepo cardRepo, config config.Config, logger *slog.Logger) (Kafka, error) {
	configSarama := sarama.NewConfig()
	configSarama.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{config.Kafka.BootstrapServers}, configSarama)
	if err != nil {
		logger.Error("Unable to create Kafka producer",
			"error", err,
			"stacktrace", stack.Trace().String())
		return &kafkaProducer{}, err
	}

	logger.Info("New Kafka connection opened")
	return &kafkaProducer{
		logger:           logger,
		producer:         producer,
		cardRepo:         cardRepo,
		bootstrapServers: config.Kafka.BootstrapServers,
		topic:            config.Kafka.ExpiredCardsTopic,
	}, nil
}

func (k *kafkaProducer) ScheduleConsumer(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Minute)
	// TODO
	// TODO
	// defer ticker.Stop() ??
	go func() {
		for {
			select {
			case <-ticker.C:
				k.logger.Info("Next tick")
				k.ProduceDeleteExpiredCards(ctx)
			case <-ctx.Done():
				k.logger.Info("Shutting down ProduceDeleteExpiredCards goroutine.")
				return
			}
		}
	}()
}

func (k *kafkaProducer) ScheduleProducer(ctx context.Context) error {
	// TODO: is this error-catch OK????
	c := cron.New()
	_, err := c.AddFunc("@every 2m", func() {
		if err := k.NewConsumer(ctx); err != nil {
			k.logger.Error("Error in NewConsumer", "error", err)
			return
		}
	})
	if err != nil {
		return err
	}
	// c.AddFunc("@every 1m", func() { fmt.Println("Every 1m") })
	c.Start()
	// TODO
	// TODO
	// defer c.Stop() ?? this stops cron as registerRepositoriesAndServices exist, may we go without it in this use case?
	return nil
}

func (k *kafkaProducer) ProduceDeleteExpiredCards(ctx context.Context) error {
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
// TODO: !!! better to create 1 consumer like producer sarama.SyncProducer ???
func (k *kafkaProducer) NewConsumer(ctx context.Context) error {
	consumer, err := sarama.NewConsumer([]string{k.bootstrapServers}, nil)
	if err != nil {
		k.logger.Error("Unable to create Kafka consumer",
			"error", err,
			"stacktrace", stack.Trace().String())
		return err
	}
	defer consumer.Close()
	k.logger.Info("Consumer started")

	partitions, err := consumer.Partitions(k.topic)
	if err != nil {
		k.logger.Error("Failed to fetch partitions",
			"topic", k.topic,
			"error", err,
			"stacktrace", stack.Trace().String())
		return err
	}

	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(k.topic, partition, sarama.OffsetNewest)

		if err != nil {
			k.logger.Error("Failed to start partition consumer",
				"partition", partition,
				"error", err,
				"stacktrace", stack.Trace().String())
			continue
		}

		// go func(pc sarama.PartitionConsumer) {
		defer partitionConsumer.Close()
		for msg := range partitionConsumer.Messages() {
			if deletedID, err := k.ConsumeCardDelete(ctx, msg); err != nil {
				k.logger.Error("Error while trying to delete", "cardID", msg.Value, "error", err)
				return err
			} else {
				k.logger.Info("Successfully deleted card ID", "deletedID", deletedID)
			}
		}
		// }(partitionConsumer)
	}
}

func (k *kafkaProducer) ConsumeCardDelete(ctx context.Context, msg *sarama.ConsumerMessage) (int64, error) {
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
