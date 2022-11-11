package order

import "time"

type Order struct {
	Pizza     string    `json:"pizza"`
	OrderedAt time.Time `json:"orderedAt"`
}

func (o *Order) GetPizza() string {
	return o.Pizza
}

func (o *Order) GetOrderedAtTime() time.Time {
	return o.OrderedAt
}
