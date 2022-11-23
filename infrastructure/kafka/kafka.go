package kafka

import (
	"factory-pattern/internal/errors"
	"factory-pattern/internal/pkg"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka interface {
	Produce(msgInBytes []byte) (kafka.Event, error)
	Consume()
}

func NewKafka(args ConfigArgs) Kafka {
	kafkaConfig := NewKafkaConfig(args)
	return &kafkaHandler{
		KafkaConfig: kafkaConfig,
		Producer:    NewProducer(kafkaConfig.GetProducerConfigMap()),
		Consumer:    NewConsumer(kafkaConfig.GetConsumerConfigMap()),
	}
}

type kafkaHandler struct {
	KafkaConfig Config
	Producer    *kafka.Producer
	Consumer    *kafka.Consumer
}

func (h *kafkaHandler) Produce(msgInBytes []byte) (kafka.Event, error) {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err := h.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     pkg.ToStringPointer(h.KafkaConfig.GetTopic()),
				Partition: kafka.PartitionAny,
			},
			Value: msgInBytes,
		},
		deliveryChan,
	)

	if err != nil {
		return nil, errors.WithStack("failed to produce message", err)
	}

	return <-deliveryChan, nil
}

func (h *kafkaHandler) Consume() {
	//TODO implement me
	panic("implement me")
}
