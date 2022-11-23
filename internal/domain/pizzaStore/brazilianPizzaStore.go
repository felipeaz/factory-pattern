package pizzaStore

import (
	"factory-pattern/infrastructure/kafka"
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
	"factory-pattern/internal/errors"
)

const (
	brazilianStoreKafkaTopic = "brazilian-store"
	brazilianConsumerGroupId = "br-consumer"
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
				KafkaConfig: kafka.ConfigArgs{
					Topic:           brazilianStoreKafkaTopic,
					BootstrapServer: "",
					ClientId:        "",
					TotalPartitions: 0,
					ConsumerConfig: kafka.ConsumerConfigArgs{
						GroupId:         brazilianConsumerGroupId,
						AutoOffsetReset: kafka.AutoOffsetResetConfig,
					},
				},
			},
		),
	}
}

func (b *brazilianPizzaStore) Order(pizza string) error {
	err := b.OrderManager.CreateOrder(pizza)
	if err != nil {
		return errors.WithStack("error calling AddToQueue", err)
	}

	return nil
}

func (b *brazilianPizzaStore) Prepare() (app.Pizza, error) {
	pizza, err := b.OrderManager.GetNextOrderInQueue()
	if err != nil {
		return nil, errors.WithStack("error calling GetNextInQueue", err)
	}

	pizza.Prepare()

	return nil, nil
}
