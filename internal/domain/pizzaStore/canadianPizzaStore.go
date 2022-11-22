package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
)

const (
	canadianStoreKafkaTopic = "canadian-store"
)

type canadianPizzaStore struct {
	PizzaFactory app.PizzaFactory
}

func NewCanadianPizzaStore() app.PizzaStore {
	return &canadianPizzaStore{
		PizzaFactory: factory.NewCanadianPizzaFactory(),
		OrderManager: order.NewManager(
			order.ManagerArgs{
				KafkaTopic: canadianStoreKafkaTopic,
			},
		),
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
