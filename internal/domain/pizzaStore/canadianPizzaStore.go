package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
	"factory-pattern/internal/errors"
)

type canadianPizzaStore struct {
	pizzaFactory app.PizzaFactory
	orderManager app.OrderManager
	worker       worker
}

func NewCanadianPizzaStore() app.PizzaStore {
	return &canadianPizzaStore{
		pizzaFactory: factory.NewCanadianPizzaFactory(),
		orderManager: order.NewManager(
			factory.KafkaHandlerFactory(factory.CanadianStore),
		),
		worker: initializeWorker(),
	}
}

func (c *canadianPizzaStore) StartWork() {
	c.worker.startWork(c.orderManager, c.pizzaFactory)
}

func (c *canadianPizzaStore) Order(pizza string) error {
	err := c.orderManager.CreateOrder(pizza)
	if err != nil {
		return errors.WithStack("error calling CreateOrder", err)
	}
	return nil
}
