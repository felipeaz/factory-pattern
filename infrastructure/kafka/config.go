package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

const (
	AutoOffsetResetConfig = "smallest"
	AcknowledgeConfig     = "acks"
)

type ConfigArgs struct {
	Topic           string
	BootstrapServer string
	ClientId        string
	TotalPartitions int
	ProducerConfig  ProducerConfigArgs
	ConsumerConfig  ConsumerConfigArgs
}

type ProducerConfigArgs struct {
	Acknowledge string
}

type ConsumerConfigArgs struct {
	GroupId         string
	AutoOffsetReset string
}

type Config interface {
	GetProducerConfigMap() *kafka.ConfigMap
	GetConsumerConfigMap() *kafka.ConfigMap
	GetTopic() string
}

func NewKafkaConfig(args ConfigArgs) Config {
	return &config{
		rawConfig: args,
	}
}

type config struct {
	producerConfigMap *kafka.ConfigMap
	consumerConfigMap *kafka.ConfigMap
	rawConfig         ConfigArgs
}

func (c *config) GetProducerConfigMap() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": c.rawConfig.BootstrapServer,
		"client.id":         c.rawConfig.ClientId,
		"acks":              c.rawConfig.ProducerConfig.Acknowledge,
	}
}

func (c *config) GetConsumerConfigMap() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": c.rawConfig.BootstrapServer,
		"client.id":         c.rawConfig.ClientId,
		"group.id":          c.rawConfig.ConsumerConfig.GroupId,
		"auto.offset.reset": c.rawConfig.ConsumerConfig.AutoOffsetReset,
	}
}

func (c *config) GetTopic() string {
	return c.rawConfig.Topic
}
