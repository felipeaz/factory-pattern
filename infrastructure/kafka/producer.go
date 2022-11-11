package kafka

import (
	"factory-pattern/infrastructure/kafka/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

// ProducerArgs arguments used in a Producer initialization
type ProducerArgs struct {
	ConfigMap *kafka.ConfigMap
}

// NewProducer returns an instance of Producer
func NewProducer(args *ProducerArgs) *kafka.Producer {
	producer, err := kafka.NewProducer(args.ConfigMap)
	if err != nil {
		log.Fatal(constants.ErrorInitializingProducer, err)
	}
	return producer
}
