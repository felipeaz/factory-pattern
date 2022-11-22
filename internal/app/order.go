package app

import (
	"log"
	"time"

	uuid "src/github.com/satori/go.uuid"
)

type Order interface {
	GetId() uuid.UUID
	GetPizza() string
	GetOrderedTime() string
}

func NewOrder(pizza string) Order {
	return &order{
		id:        newUUID(),
		pizza:     pizza,
		orderedAt: time.Now().UTC(),
	}
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

func newUUID() uuid.UUID {
	orderUUID, err := uuid.NewV4()
	if err != nil {
		log.Fatal("failed to initialize order UUID")
	}
	return orderUUID
}
