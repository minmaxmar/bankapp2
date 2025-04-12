package kafka

import (
	"bankapp2/app/models"
	"bankapp2/helper/config"
	"context"
	"strconv"

	"log/slog"

	"github.com/go-stack/stack"
	"gorm.io/gorm"

	"github.com/IBM/sarama"
)

var (
	topicDeleteCard string = "delete_card"
)

type kafkaProducer struct {
	logger   *slog.Logger
	producer sarama.SyncProducer
	cardRepo cardRepo
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
	NewConsumer(ctx context.Context, config config.Config) error
	ConsumeCardDelete(ctx context.Context, msg *sarama.ConsumerMessage) (int64, error)
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
		logger:   logger,
		producer: producer,
		cardRepo: cardRepo,
	}, nil
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
			Topic: topicDeleteCard,
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
func (k *kafkaProducer) NewConsumer(ctx context.Context, config config.Config) error {
	consumer, err := sarama.NewConsumer([]string{config.Kafka.BootstrapServers}, nil)
	if err != nil {
		k.logger.Error("Unable to create Kafka consumer",
			"error", err,
			"stacktrace", stack.Trace().String())
		return err
	}
	defer consumer.Close()
	k.logger.Info("Consumer started")

	partitions, err := consumer.Partitions(topicDeleteCard)
	if err != nil {
		k.logger.Error("Failed to fetch partitions",
			"topic", topicDeleteCard,
			"error", err,
			"stacktrace", stack.Trace().String())
		return err
	}

	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topicDeleteCard, partition, sarama.OffsetNewest)

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
