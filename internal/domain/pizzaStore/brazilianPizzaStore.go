package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
	"factory-pattern/internal/errors"
)

const (
	brazilianStoreKafkaTopic = "brazilian-store"
)

type brazilianPizzaStore struct {
	PizzaFactory app.PizzaFactory
	OrderManager app.OrderManager
}

func NewBrazilianPizzaStore() app.PizzaStore {
	return &brazilianPizzaStore{
		PizzaFactory: factory.NewBrazilianPizzaFactory(),
		OrderManager: order.NewManager(
			order.ManagerArgs{
				KafkaTopic: brazilianStoreKafkaTopic,
			},
		),
	}
}

func (b *brazilianPizzaStore) Order(pizza string) (int, error) {
	orderedPizza, err := b.PizzaFactory.CreatePizza(pizza)
	if err != nil {
		return 0, errors.WithStack("error calling CreatePizza", err)
	}

	positionInQueue, err := b.OrderManager.AddToQueue(orderedPizza)
	if err != nil {
		return 0, errors.WithStack("error calling AddToQueue", err)
	}

	return positionInQueue, nil
}

func (b *brazilianPizzaStore) Prepare(order app.OrderManager) (*app.Pizza, error) {
	pizza, err := b.OrderManager.GetNextInQueue()
	if err != nil {
		return nil, errors.WithStack("error calling GetNextInQueue", err)
	}
	pizza.Prepare()
	return nil, nil
}
