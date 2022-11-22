package order

import (
	"factory-pattern/infrastructure/kafka"
	"factory-pattern/internal/app"
)

type ManagerArgs struct {
	KafkaTopic  string
	KafkaConfig kafka.ConfigArgs
}

type orderManager struct {
	KafkaHandler kafka.Kafka
}

func NewManager(args ManagerArgs) app.OrderManager {
	return &orderManager{
		KafkaHandler: kafka.NewKafka(args.KafkaConfig),
	}
}

func (o *orderManager) AddToQueue(pizza app.Pizza) (positionInQueue int, err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderManager) GetNextInQueue() (pizza app.Pizza, err error) {
	//TODO implement me
	panic("implement me")
}
