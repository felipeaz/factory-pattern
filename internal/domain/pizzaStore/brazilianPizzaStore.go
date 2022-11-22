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

func (b *brazilianPizzaStore) Order(pizza string) (app.Order, error) {
	orderObj, err := b.OrderManager.CreateOrder(pizza)
	if err != nil {
		return nil, errors.WithStack("error calling AddToQueue", err)
	}

	return orderObj, nil
}

func (b *brazilianPizzaStore) Prepare() (app.Pizza, error) {
	pizza, err := b.OrderManager.GetNextOrderInQueue()
	if err != nil {
		return nil, errors.WithStack("error calling GetNextInQueue", err)
	}
	pizza.Prepare()
	return nil, nil
}
