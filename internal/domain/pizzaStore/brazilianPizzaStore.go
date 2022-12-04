package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
	"factory-pattern/internal/errors"
)

type brazilianPizzaStore struct {
	pizzaFactory app.PizzaFactory
	orderManager app.OrderManager
	worker       worker
}

func NewBrazilianPizzaStore() app.PizzaStore {
	return &brazilianPizzaStore{
		pizzaFactory: factory.NewBrazilianPizzaFactory(),
		orderManager: order.NewManager(
			factory.KafkaHandlerFactory(factory.BrazilianStore),
		),
		worker: initializeWorker(),
	}
}

func (b *brazilianPizzaStore) StartWork() {
	b.worker.startWork(b.orderManager, b.pizzaFactory)
}

func (b *brazilianPizzaStore) Order(pizza string) error {
	err := b.orderManager.CreateOrder(pizza)
	if err != nil {
		return errors.WithStack("error calling CreateOrder", err)
	}
	return nil
}
