package event

import (
	"factory-pattern/internal/errors"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func HandleEvent(event kafka.Event) error {
	switch eventType := event.(type) {
	case *kafka.Message:
		return handleMessage(event, eventType)
	default:
		return handleDefaultType(eventType)
	}
}

func handleMessage(event kafka.Event, message *kafka.Message) error {
	if err := message.TopicPartition.Error; err != nil {
		return err
	}
	log.Println("Event registered", event)
	return nil
}

func handleDefaultType(eventType any) error {
	log.Println("Event type not mapped", eventType)
	return errors.New("event type not mapped")
}
