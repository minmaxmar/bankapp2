package kafka

import (
	"bankapp2/helper/config"
	"context"
	"fmt"
	"os"

	"log/slog"

	"github.com/go-stack/stack"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	topicDeleteCard string = "delete_card"
)

type kafkaProducer struct {
	logger *slog.Logger

	producer *kafka.Producer

	cardRepo cardRepo
}

type cardRepo interface {
	DeleteCardID(ctx context.Context, id int) (int64, error)
	GetExpiredCards(ctx context.Context) (int64, error) // TODO: implement in cards repo
}

type Kafka interface {
	// ProduceDeleteCard(ctx context.Context, id int) error
	ProduceDeleteExpiredCards(ctx context.Context) error
	NewConsumer(ctx context.Context, config config.Config)
	ConsumeCardDelete(ctx context.Context, msg *kafka.Message)
}

func NewConn(cardRepo cardRepo, config config.Config, logger *slog.Logger) Kafka {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.BootstrapServers,
	})
	if err != nil {
		logger.Error("Unable to create Kafka producer",
			"error", err,
			"stacktrace", stack.Trace().String())
		os.Exit(1)
	}

	logger.Info("New Kafka connection opened")
	return &kafkaProducer{
		logger:   logger,
		producer: producer,
		cardRepo: cardRepo,
	}
}

func (k *kafkaProducer) ProduceDeleteExpiredCards(ctx context.Context) error {
	// TODO: select from card where expire_date >= now
	// k.cardRepo.GetExpiredCards()
	id := 1
	message := fmt.Sprintf("%d", id) // serialize message
	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topicDeleteCard, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		k.logger.Error("Failed to produce message", "error", err)
		return err
	}

	k.producer.Flush(1000)
	return nil
}

func (k *kafkaProducer) NewConsumer(ctx context.Context, config config.Config) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.BootstrapServers,
		"group.id":          "card_delete_group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		k.logger.Error("Unable to create Kafka consumer",
			"error", err,
			"stacktrace", stack.Trace().String())
		os.Exit(1)
	}

	err = consumer.Subscribe(topicDeleteCard, nil)
	if err != nil {
		k.logger.Error("Failed to subscribe to topic",
			"topic", topicDeleteCard,
			"error", err,
			"stacktrace", stack.Trace().String())
		os.Exit(1)
	}

	k.logger.Info("Consumer started")

	go func() {
		for {
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				k.ConsumeCardDelete(ctx, msg)
			} else {
				k.logger.Error("Error while consuming message", "error", err)
			}
		}
	}()
}

func (k *kafkaProducer) ConsumeCardDelete(ctx context.Context, msg *kafka.Message) {
	cardID := string(msg.Value)
	k.logger.Info("Received message to delete card ID: %s\n", cardID)
	// convert cardID to an integer and k.cardRepo.DeleteCardID
}
