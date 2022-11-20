package order

import (
	"factory-pattern/internal/app"
)

type orderManager struct {
	topic string
}

func NewOrderManager(kafkaTopic string) app.OrderManager {
	return &orderManager{
		topic: kafkaTopic,
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
