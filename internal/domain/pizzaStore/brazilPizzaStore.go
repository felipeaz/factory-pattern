package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
)

type brazilianPizzaStore struct {
	PizzaFactory app.PizzaFactory
}

func NewBrazilianPizzaStore() app.PizzaStore {
	return &brazilianPizzaStore{
		PizzaFactory: factory.NewBrazilianPizzaFactory(),
	}
}

func (b *brazilianPizzaStore) Order(pizza string) {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianPizzaStore) Prepare(order app.Order) *app.Pizza {
	//TODO implement me
	panic("implement me")
}
