package event

import (
	"factory-pattern/internal/errors"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func HandleEvent(event kafka.Event) (*kafka.Message, error) {
	switch eventType := event.(type) {
	case *kafka.Message:
		return handleMessage(event, eventType)
	case kafka.Error:
		return nil, handleKafkaError(event)
	default:
		return nil, handleDefaultType(eventType)
	}
}

func handleMessage(event kafka.Event, kafkaMessage *kafka.Message) (*kafka.Message, error) {
	if err := kafkaMessage.TopicPartition.Error; err != nil {
		return kafkaMessage, err
	}
	log.Println("Event registered", event)
	return kafkaMessage, nil
}

func handleKafkaError(event any) error {
	log.Println("Event error", event)
	return errors.New("event error:", event)
}

func handleDefaultType(event any) error {
	log.Println("Event type not mapped, ignoring", event)
	return nil
}
