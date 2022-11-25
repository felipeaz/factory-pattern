package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
)

type nyPizzaStore struct {
	PizzaFactory app.PizzaFactory
	OrderManager app.OrderManager
}

func NewNYPizzaStore() app.PizzaStore {
	return &nyPizzaStore{
		PizzaFactory: factory.NewNYPizzaFactory(),
		OrderManager: order.NewManager(
			factory.KafkaHandlerFactory(factory.NewYorkStore),
		),
	}
}

func (n *nyPizzaStore) StartWork() {

}

func (n *nyPizzaStore) Order(pizza string) error {
	//TODO implement me
	panic("implement me")
}

func (n *nyPizzaStore) prepare() (app.Pizza, error) {
	//TODO implement me
	panic("implement me")
}

func (n *nyPizzaStore) delivery(pizza app.Pizza) bool {
	return true
}
