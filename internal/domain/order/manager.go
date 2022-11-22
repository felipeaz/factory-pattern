package order

import (
	"factory-pattern/infrastructure/kafka"
	"factory-pattern/internal/app"
	"factory-pattern/internal/errors"
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

func (o *orderManager) CreateOrder(pizza string) (order app.Order, err error) {
	order = app.NewOrder(pizza)

	err = o.placeOrderInQueue(order)
	if err != nil {
		return nil, errors.WithStack("failed to place an order", err)
	}

	return order, nil
}

func (o *orderManager) GetNextOrderInQueue() (pizza app.Pizza, err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderManager) placeOrderInQueue(order app.Order) error {
	//TODO implement me
	panic("implement me")
}
