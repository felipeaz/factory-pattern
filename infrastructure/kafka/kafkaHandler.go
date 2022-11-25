package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"

	"factory-pattern/infrastructure/kafka/event"
	"factory-pattern/internal/errors"
	"factory-pattern/internal/pkg"
)

type Kafka interface {
	Produce(msgInBytes []byte) (kafka.Event, error)
	Consume(ordersEventCh chan<- *kafka.Message, done <-chan bool)
}

func NewKafka(args ConfigArgs) Kafka {
	kafkaConfig := NewKafkaConfig(args)
	return &kafkaHandler{
		KafkaConfig: kafkaConfig,
		Producer:    NewProducer(kafkaConfig.GetProducerConfigMap()),
		Consumer:    NewConsumer(kafkaConfig.GetConsumerConfigMap(), kafkaConfig.GetTopic()),
	}
}

type kafkaHandler struct {
	KafkaConfig Config
	Producer    *kafka.Producer
	Consumer    *kafka.Consumer
}

func (h *kafkaHandler) Produce(msgInBytes []byte) (kafka.Event, error) {
	produceChan := make(chan kafka.Event)
	defer close(produceChan)

	err := h.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     pkg.ToStringPointer(h.KafkaConfig.GetTopic()),
				Partition: kafka.PartitionAny,
			},
			Value: msgInBytes,
		},
		produceChan,
	)

	if err != nil {
		return nil, errors.WithStack("failed to produce message", err)
	}

	return <-produceChan, nil
}

func (h *kafkaHandler) Consume(ordersEventCh chan<- *kafka.Message, done <-chan bool) {
	defer h.Consumer.Close()

	select {
	case <-done:
		log.Println("Consumer stopped pooling")
		return
	default:
		consumerEvent := h.Consumer.Poll(100)

		msg, err := event.HandleEvent(consumerEvent)
		if err != nil {
			return
		}

		ordersEventCh <- msg
	}
}
