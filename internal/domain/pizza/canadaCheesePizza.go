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

func (p *canadianCheesePizza) Prepare() bool {
	//TODO implement me
	panic("implement me")
}

func (p *canadianCheesePizza) Cut() {
	//TODO implement me
	panic("implement me")
}

func (p *canadianCheesePizza) Box() {
	//TODO implement me
	panic("implement me")
}

func (p *canadianCheesePizza) GetName() string {
	return p.Name
}

func (p *canadianCheesePizza) GetDescription() string {
	return p.Description
}
