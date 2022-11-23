package pizzaStore

import (
	"factory-pattern/infrastructure/kafka"
	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/factory"
	"factory-pattern/internal/domain/order"
)

const (
	canadianStoreKafkaTopic      = "canadian-store"
	canadianStoreConsumerGroupId = "ca-consumer"
)

type canadianPizzaStore struct {
	PizzaFactory app.PizzaFactory
	OrderManager app.OrderManager
}

func NewCanadianPizzaStore() app.PizzaStore {
	return &canadianPizzaStore{
		PizzaFactory: factory.NewCanadianPizzaFactory(),
		OrderManager: order.NewManager(
			order.ManagerArgs{
				KafkaConfig: kafka.ConfigArgs{
					Topic:           canadianStoreKafkaTopic,
					BootstrapServer: "",
					ClientId:        "",
					TotalPartitions: 0,
					ConsumerConfig: kafka.ConsumerConfigArgs{
						GroupId:         canadianStoreConsumerGroupId,
						AutoOffsetReset: kafka.AutoOffsetResetConfig,
					},
				},
			},
		),
	}
}

func (c *canadianPizzaStore) Order(pizza string) error {
	//TODO implement me
	panic("implement me")
}

func (c *canadianPizzaStore) Prepare() (app.Pizza, error) {
	//TODO implement me
	panic("implement me")
}
