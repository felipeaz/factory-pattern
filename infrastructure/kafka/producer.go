package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"factory-pattern/infrastructure/kafka/constants"
)

type ProducerArgs struct {
	ConfigMap *kafka.ConfigMap
}

func NewProducer(args ProducerArgs) *kafka.Producer {
	producer, err := kafka.NewProducer(args.ConfigMap)
	if err != nil {
		log.Fatal(constants.ErrorInitializingProducer, err)
	}
	return producer
}
