package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
)

const (
	canadianStoreKafkaTopic = "canadian-store"
)

type canadianPizzaStore struct {
	PizzaFactory app.PizzaFactory
	OrderManager app.OrderManager
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

func (c *canadianPizzaStore) Order(pizza string) (app.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (c *canadianPizzaStore) Prepare() (app.Pizza, error) {
	//TODO implement me
	panic("implement me")
}
