package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

const (
	AutoOffsetResetConfig = "smallest"
)

type ConfigArgs struct {
	Topic           string
	BootstrapServer string
	ClientId        string
	TotalPartitions int
	ConsumerConfig  ConsumerConfigArgs
}

type ConsumerConfigArgs struct {
	GroupId         string
	AutoOffsetReset string
}

type Config interface {
	GetConfigMap() *kafka.ConfigMap
}

func NewBaseConfig(args ConfigArgs) Config {
	return &config{
		configMap: buildConfigMap(args),
	}
}

type config struct {
	configMap *kafka.ConfigMap
}

func (c *config) GetConfigMap() *kafka.ConfigMap {
	return c.configMap
}

func buildConfigMap(args ConfigArgs) *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": args.BootstrapServer,
		"client.id":         args.ClientId,
		"acks":              "all",
	}
}
