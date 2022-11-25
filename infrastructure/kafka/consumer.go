package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"factory-pattern/infrastructure/kafka/constants"
)

func NewConsumer(configMap *kafka.ConfigMap, topics string) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal(constants.ErrorInitializingConsumer)
	}
	err = consumer.Subscribe(topics, nil)
	if err != nil {
		log.Fatal(constants.ErrorSubscribingToTopics)
	}
	return consumer
}
