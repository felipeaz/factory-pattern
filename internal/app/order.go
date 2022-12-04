package app

import (
	uuid "github.com/satori/go.uuid"
)

type Order interface {
	GetId() uuid.UUID
	GetPizza() string
	GetOrderedTime() string
	ToBytes() ([]byte, error)
}
