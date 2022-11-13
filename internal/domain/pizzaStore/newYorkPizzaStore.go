package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
)

type nyPizzaStore struct {
	PizzaFactory app.PizzaFactory
}

func NewNYPizzaStore() app.PizzaStore {
	return &nyPizzaStore{
		PizzaFactory: factory.NewNYPizzaFactory(),
	}
}

func (n *nyPizzaStore) Order(pizza string) {
	//TODO implement me
	panic("implement me")
}

func (n *nyPizzaStore) Prepare(order app.Order) *app.Pizza {
	//TODO implement me
	panic("implement me")
}
