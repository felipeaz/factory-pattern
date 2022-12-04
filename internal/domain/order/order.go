package order

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"

	"factory-pattern/internal/app"
	"factory-pattern/internal/errors"
)

func NewOrderFromPizzaName(pizza string) app.Order {
	return &order{
		id:        uuid.NewV4(),
		pizza:     pizza,
		orderedAt: time.Now().UTC(),
	}
}

func NewOrderInBytes(pizza string) ([]byte, error) {
	return NewOrderFromPizzaName(pizza).ToBytes()
}

type order struct {
	id        uuid.UUID
	pizza     string
	orderedAt time.Time
}

func (o *order) GetId() uuid.UUID {
	return o.id
}

func (o *order) GetPizza() string {
	return o.pizza
}

func (o *order) GetOrderedTime() string {
	return o.orderedAt.Format("15:04:05")
}

func (o *order) ToBytes() ([]byte, error) {
	dto := o.getOrderDTO()

	b, err := json.Marshal(dto)
	if err != nil {
		return nil, errors.WithStack("failed to marshal order object", err)
	}

	return b, nil
}

func (o *order) getOrderDTO() OrderDto {
	return OrderDto{
		Id:        o.GetId().String(),
		Pizza:     o.GetPizza(),
		OrderedAt: o.GetOrderedTime(),
	}
}
