package order

import (
	"factory-pattern/infrastructure/kafka"
	"factory-pattern/infrastructure/kafka/event"
	"factory-pattern/internal/app"
	"factory-pattern/internal/errors"
)

type orderManager struct {
	kafkaHandler kafka.Kafka
	topic        string
	groupId      string
}

func NewManager(kafkaHandler kafka.Kafka) app.OrderManager {
	return &orderManager{
		kafkaHandler: kafkaHandler,
	}
}

func (o *orderManager) CreateOrder(pizza string) (err error) {
	orderedPizza, err := NewOrderInBytes(pizza)
	if err != nil {
		return errors.WithStack("failed to create order in bytes", err)
	}

	err = o.placeOrderInQueue(orderedPizza)
	if err != nil {
		return errors.WithStack("failed to place an order in queue", err)
	}

	return err
}

func (o *orderManager) placeOrderInQueue(order []byte) error {
	kafkaEvent, err := o.kafkaHandler.Produce(order)
	if err != nil {
		return errors.WithStack("error calling producer", err)
	}

	return event.HandleEvent(kafkaEvent)
}

func (o *orderManager) GetNextOrderInQueue() (pizza app.Pizza, err error) {
	//TODO implement me
	panic("implement me")
}
