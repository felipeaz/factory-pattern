package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
)

const (
	newYorkStoreKafkaTopic = "new-york-store"
)

type nyPizzaStore struct {
	PizzaFactory app.PizzaFactory
	OrderManager app.OrderManager
}

func NewNYPizzaStore() app.PizzaStore {
	return &nyPizzaStore{
		PizzaFactory: factory.NewNYPizzaFactory(),
		OrderManager: order.NewManager(
			order.ManagerArgs{
				KafkaTopic: newYorkStoreKafkaTopic,
			},
		),
	}
}

func (n *nyPizzaStore) Order(pizza string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (n *nyPizzaStore) Prepare(order app.OrderManager) (*app.Pizza, error) {
	//TODO implement me
	panic("implement me")
}
