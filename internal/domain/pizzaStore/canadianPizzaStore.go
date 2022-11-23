package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
)

type canadianPizzaStore struct {
	PizzaFactory app.PizzaFactory
	OrderManager app.OrderManager
}

func NewCanadianPizzaStore() app.PizzaStore {
	return &canadianPizzaStore{
		PizzaFactory: factory.NewCanadianPizzaFactory(),
		OrderManager: order.NewManager(
			factory.KafkaHandlerFactory(factory.CanadianStore),
		),
	}
}

func (c *canadianPizzaStore) Order(pizza string) error {
	//TODO implement me
	panic("implement me")
}

func (c *canadianPizzaStore) Prepare() (app.Pizza, error) {
	//TODO implement me
	panic("implement me")
}
