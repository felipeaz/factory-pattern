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
	done := make(chan bool)
	defer close(done)

	pizzaQueueChan := make(chan string)
	defer close(pizzaQueueChan)

	go b.OrderManager.GetNextOrderInQueue(pizzaQueueChan, done)

	for orderedPizza := range pizzaQueueChan {
		pizzaInProduction, err := b.PizzaFactory.CreatePizza(orderedPizza)
		if err != nil {
			done <- true
		}

		select {
		case <-pizzaInProduction.Prepare():
			log.Printf("%s Pizza is ready for delivery", pizzaInProduction.GetName())

			pizzaInProduction.Cut()
			pizzaInProduction.Box()
			delivery(pizzaInProduction)

			continue
		default:
			log.Printf("Preparing a %s Pizza", pizzaInProduction.GetName())
		}
	}
}

func (b *brazilianPizzaStore) Order(pizza string) error {
	err := b.OrderManager.CreateOrder(pizza)
	if err != nil {
		return errors.WithStack("error calling CreateOrder", err)
	}
	return nil
}

func delivery(pizza app.Pizza) {
	log.Printf("A %s Pizza is out for delivery", pizza.GetDescription())
}
