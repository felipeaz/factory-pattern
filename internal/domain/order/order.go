package order

import (
	"factory-pattern/internal/app"
	"time"
)

type order struct {
	Pizza     string    `json:"pizza"`
	OrderedAt time.Time `json:"orderedAt"`
}

func NewOrder(pizza string) app.Order {
	return &order{
		Pizza:     pizza,
		OrderedAt: time.Now().UTC(),
	}
}

func (o *order) GetPizza() string {
	return o.Pizza
}

func (o *order) GetOrderedAtTime() time.Time {
	return o.OrderedAt
}
