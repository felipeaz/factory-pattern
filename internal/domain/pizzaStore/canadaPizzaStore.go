package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
)

type canadianPizzaStore struct {
	PizzaFactory app.PizzaFactory
}

func NewCanadianPizzaStore() app.PizzaStore {
	return &canadianPizzaStore{
		PizzaFactory: factory.NewCanadianPizzaFactory(),
	}
}

func (c *canadianPizzaStore) Order(pizza string) {
	//TODO implement me
	panic("implement me")
}

func (c *canadianPizzaStore) Prepare(order app.Order) *app.Pizza {
	//TODO implement me
	panic("implement me")
}
