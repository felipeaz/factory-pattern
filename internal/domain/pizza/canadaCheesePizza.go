package pizza

import (
	"time"

	"factory-pattern/internal/app"
)

const (
	CanadianCheesePizzaTitle       = "Cheese"
	CanadianCheesePizzaDescription = "Delicious Canadian Style Cheese Pizza"
)

type canadianCheesePizza struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewCanadianCheesePizza() app.Pizza {
	return &canadianCheesePizza{
		Name:               CanadianCheesePizzaTitle,
		Description:        CanadianCheesePizzaDescription,
		TimeToPrepareInSec: time.Second * 2,
	}
}

func (b *canadianCheesePizza) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (b *canadianCheesePizza) Cut() {
	//TODO implement me
	panic("implement me")
}

func (b *canadianCheesePizza) Box() {
	//TODO implement me
	panic("implement me")
}

func (b *canadianCheesePizza) IsPrepared() bool {
	//TODO implement me
	panic("implement me")
}
