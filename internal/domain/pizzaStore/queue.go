package pizzaStore

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"

	queueManager "factory-pattern/infrastructure/kafka"
)

func buildQueueManager() *kafka.Producer {
	args := &queueManager.ProducerArgs{
		ConfigMap: &kafka.ConfigMap{},
	}
	return queueManager.NewProducer(args)
}
