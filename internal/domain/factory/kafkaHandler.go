package factory

import (
	"factory-pattern/infrastructure/kafka"
	"log"
)

const (
	BrazilianStore = "brazilian"
	CanadianStore  = "canadian"
	NewYorkStore   = "newYork"

	brazilianStoreKafkaTopic = "brazilian-store"
	brazilianConsumerGroupId = "br-consumer"

	canadianStoreKafkaTopic      = "canadian-store"
	canadianStoreConsumerGroupId = "ca-consumer"

	newYorkStoreKafkaTopic      = "new-york-store"
	newYorkStoreConsumerGroupId = "ny-consumer"
)

func KafkaHandlerFactory(store string) kafka.Kafka {
	switch store {
	case BrazilianStore:
		return buildKafkaHandler(brazilianStoreKafkaTopic, brazilianConsumerGroupId)
	case CanadianStore:
		return buildKafkaHandler(canadianStoreKafkaTopic, canadianStoreConsumerGroupId)
	case NewYorkStore:
		return buildKafkaHandler(newYorkStoreKafkaTopic, newYorkStoreConsumerGroupId)
	default:
		log.Fatal("store does not exist")
		return nil
	}
}

func buildKafkaHandler(topic, groupId string) kafka.Kafka {
	return kafka.NewKafka(
		kafka.ConfigArgs{
			Topic:           topic,
			BootstrapServer: "localhost:9092",
			ClientId:        "localhost",
			TotalPartitions: 0,
			ProducerConfig: kafka.ProducerConfigArgs{
				Acknowledge: kafka.AcknowledgeConfig,
			},
			ConsumerConfig: kafka.ConsumerConfigArgs{
				GroupId:         groupId,
				AutoOffsetReset: kafka.AutoOffsetResetConfig,
			},
		},
	)
}
