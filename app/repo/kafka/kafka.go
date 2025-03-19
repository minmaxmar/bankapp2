package kafka

import (
	"bankapp2/helper/config"
	"context"
	"fmt"

	"log/slog"

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
	// DeleteCardID(ctx context.Context, id int) (int64, error)
	GetExpiredCards(ctx context.Context) (int64, error) // TODO: implement in cards repo
}

type Kafka interface {
	// ProduceDeleteCard(ctx context.Context, id int) error
	ProduceDeleteExpiredCards(ctx context.Context) error
	NewConsumer(ctx context.Context)
	ConsumeCardDelete(ctx context.Context, msg *kafka.Message)
}

func NewConn(cardRepo cardRepo, config config.Config, logger *slog.Logger) Kafka {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.BootstrapServers,
	})
	if err != nil {
		logger.Fatalf("Unable to create Kafka producer: %v\n", err)
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

	message := fmt.Sprintf("%d", id) // serialize message
	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topicDeleteCard, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		k.logger.Printf("Failed to produce message: %v\n", err)
		return err
	}

	k.producer.Flush(1000)
	return nil
}

func (k *kafkaProducer) NewConsumer(ctx context.Context) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.BootstrapServers,
		"group.id":          "card_delete_group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		k.logger.Fatalf("Failed to create Kafka consumer: %v\n", err)
	}

	err = consumer.Subscribe(topicDeleteCard, nil)
	if err != nil {
		k.logger.Fatalf("Failed to subscribe to topic %s: %v\n", topicDeleteCard, err)
	}

	k.logger.Println("Consumer started")

	// Consuming messages in a goroutine
	go func() {
		for {
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				k.ConsumeCardDelete(ctx, msg)
			} else {
				k.logger.Printf("Error while consuming message: %v\n", err)
			}
		}
	}()
}

func (k *kafkaProducer) ConsumeCardDelete(ctx context.Context, msg *kafka.Message) {
	cardID := string(msg.Value)
	// Add logic to handle the card ID (delete it using cardRepo)
	k.logger.Printf("Received message to delete card ID: %s\n", cardID)
	// Here you might want to convert cardID to an integer and call cardRepo.DeleteCardID
}
