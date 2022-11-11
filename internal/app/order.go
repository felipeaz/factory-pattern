package app

import "time"

type Order interface {
	GetPizza() string
	GetOrderedAtTime() time.Time
}
