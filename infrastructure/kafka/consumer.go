package kafka

import (
	"factory-pattern/infrastructure/kafka/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

// ConsumerArgs arguments used in a Consumer initialization
type ConsumerArgs struct {
	ConfigMap *kafka.ConfigMap
}

// NewConsumer returns an instance of Consumer
func NewConsumer(args *ConsumerArgs) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(args.ConfigMap)
	if err != nil {
		log.Fatal(constants.ErrorInitializingConsumer)
	}
	return consumer
}
