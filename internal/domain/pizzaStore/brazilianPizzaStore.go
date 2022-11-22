package pizzaStore

import (
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
)

const (
	brazilianStoreKafkaTopic = "brazilian-store"
)

type brazilianPizzaStore struct {
	PizzaFactory app.PizzaFactory
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

func (b *brazilianPizzaStore) Order(pizza string) {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianPizzaStore) Prepare(order app.Order) *app.Pizza {
	//TODO implement me
	panic("implement me")
}
