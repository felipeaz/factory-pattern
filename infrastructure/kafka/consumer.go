package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"factory-pattern/infrastructure/kafka/constants"
)

type ConsumerArgs struct {
	ConfigMap *kafka.ConfigMap
}

func NewConsumer(args ConsumerArgs) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(args.ConfigMap)
	if err != nil {
		log.Fatal(constants.ErrorInitializingConsumer)
	}
	return consumer
}
