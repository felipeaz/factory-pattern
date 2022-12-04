package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
	"factory-pattern/internal/errors"
)

type nyPizzaStore struct {
	pizzaFactory app.PizzaFactory
	orderManager app.OrderManager
	worker       worker
}

func NewNYPizzaStore() app.PizzaStore {
	return &nyPizzaStore{
		pizzaFactory: factory.NewNYPizzaFactory(),
		orderManager: order.NewManager(
			factory.KafkaHandlerFactory(factory.NewYorkStore),
		),
		worker: initializeWorker(),
	}
}

func (n *nyPizzaStore) StartWork() {
	n.worker.startWork(n.orderManager, n.pizzaFactory)
}

func (n *nyPizzaStore) Order(pizza string) error {
	err := n.orderManager.CreateOrder(pizza)
	if err != nil {
		return errors.WithStack("error calling CreateOrder", err)
	}
	return nil
}
