package pizza

import (
	"time"

	"factory-pattern/internal/app"
)

const (
	BrazilianCheesePizzaTitle       = "Cheese"
	BrazilianCheesePizzaDescription = "Delicious Brazilian Style Cheese Pizza"
)

type brazilianCheesePizza struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewBrazilianCheesePizza() app.Pizza {
	return &brazilianCheesePizza{
		Name:               BrazilianCheesePizzaTitle,
		Description:        BrazilianCheesePizzaDescription,
		TimeToPrepareInSec: time.Second * 4,
	}
}

func (b *brazilianCheesePizza) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianCheesePizza) Cut() {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianCheesePizza) Box() {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianCheesePizza) GetTimeToPrepareInSeconds() time.Duration {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianCheesePizza) IsPrepared() bool {
	//TODO implement me
	panic("implement me")
}
