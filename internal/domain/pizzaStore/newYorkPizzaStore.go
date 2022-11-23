package pizzaStore

import (
	"factory-pattern/infrastructure/kafka"
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
)

const (
	newYorkStoreKafkaTopic      = "new-york-store"
	newYorkStoreConsumerGroupId = "ny-consumer"
)

type nyPizzaStore struct {
	PizzaFactory app.PizzaFactory
	OrderManager app.OrderManager
}

func NewNYPizzaStore() app.PizzaStore {
	return &nyPizzaStore{
		PizzaFactory: factory.NewNYPizzaFactory(),
		OrderManager: order.NewManager(
			order.ManagerArgs{
				KafkaConfig: kafka.ConfigArgs{
					Topic:           newYorkStoreKafkaTopic,
					BootstrapServer: "",
					ClientId:        "",
					TotalPartitions: 0,
					ConsumerConfig: kafka.ConsumerConfigArgs{
						GroupId:         newYorkStoreConsumerGroupId,
						AutoOffsetReset: kafka.AutoOffsetResetConfig,
					},
				},
			},
		),
	}
}

func (n *nyPizzaStore) Order(pizza string) error {
	//TODO implement me
	panic("implement me")
}

func (n *nyPizzaStore) Prepare() (app.Pizza, error) {
	//TODO implement me
	panic("implement me")
}
