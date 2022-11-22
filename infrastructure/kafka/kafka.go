package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Kafka interface {
	Produce()
	Consume()
}

func NewKafka(args ConfigArgs) Kafka {
	kafkaConfig := NewBaseConfig(args)
	return &handler{
		Producer: NewProducer(
			ProducerArgs{
				ConfigMap: kafkaConfig.GetConfigMap(),
			},
		),
		Consumer: NewConsumer(
			ConsumerArgs{
				ConfigMap: kafkaConfig.GetConfigMap(),
			},
		),
	}
}

type handler struct {
	Producer *kafka.Producer
	Consumer *kafka.Consumer
}

func (h handler) Produce() {
	//TODO implement me
	panic("implement me")
}

func (h handler) Consume() {
	//TODO implement me
	panic("implement me")
}
