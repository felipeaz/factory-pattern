package pizzaStore

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
)

const (
	brazilianStoreKafkaTopic = "brazilian-store"
)

type brazilianPizzaStore struct {
	PizzaFactory app.PizzaFactory
	QueueManager *kafka.Producer
}

func NewBrazilianPizzaStore() app.PizzaStore {
	return &brazilianPizzaStore{
		PizzaFactory: factory.NewBrazilianPizzaFactory(),
		QueueManager: buildQueueManager(),
	}
}

func (b *brazilianPizzaStore) Order(pizza string) {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianPizzaStore) Prepare(order app.Order) *app.Pizza {
	//TODO implement me
	panic("implement me")
}
