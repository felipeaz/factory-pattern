package pizzaStore

import (
	"log"

	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
	"factory-pattern/internal/errors"
)

type brazilianPizzaStore struct {
	PizzaFactory app.PizzaFactory
	OrderManager app.OrderManager
}

func NewBrazilianPizzaStore() app.PizzaStore {
	return &brazilianPizzaStore{
		PizzaFactory: factory.NewBrazilianPizzaFactory(),
		OrderManager: order.NewManager(
			factory.KafkaHandlerFactory(factory.BrazilianStore),
		),
	}
}

func (b *brazilianPizzaStore) StartWork() {

}

func (b *brazilianPizzaStore) Order(pizza string) error {
	err := b.OrderManager.CreateOrder(pizza)
	if err != nil {
		return errors.WithStack("error calling CreateOrder", err)
	}

	return nil
}

func (b *brazilianPizzaStore) prepare() (app.Pizza, error) {
	doneCh := make(chan bool)
	pizza, err := b.OrderManager.GetNextOrderInQueue()
	if err != nil {
		return nil, errors.WithStack("error calling GetNextInQueue", err)
	}

	doneCh <- pizza.Prepare()
	select {
	case <-doneCh:
		log.Printf("%s Pizza is ready for delivery\n", pizza.GetName())
	default:
		log.Printf("Preparing a %s Pizza\n", pizza.GetName())
	}

	return pizza, nil
}
func (b *brazilianPizzaStore) delivery(pizza app.Pizza) bool {
	log.Printf("A %s Pizza is out for delivery", pizza.GetDescription())
	return true
}
