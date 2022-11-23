package order

import (
	"factory-pattern/infrastructure/kafka"
	"factory-pattern/infrastructure/kafka/event"
	"factory-pattern/internal/app"
	"factory-pattern/internal/errors"
)

type ManagerArgs struct {
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
	kafkaEvent, err := o.KafkaHandler.Produce(order)
	if err != nil {
		return errors.WithStack("error calling producer", err)
	}

	return event.HandleEvent(kafkaEvent)
}

func (o *orderManager) GetNextOrderInQueue() (pizza app.Pizza, err error) {
	//TODO implement me
	panic("implement me")
}
