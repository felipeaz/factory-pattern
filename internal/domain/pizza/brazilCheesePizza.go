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

func (p *brazilianCheesePizza) Prepare() chan bool {
	//TODO implement me
	panic("implement me")
}

func (p *brazilianCheesePizza) Cut() {
	//TODO implement me
	panic("implement me")
}

func (p *brazilianCheesePizza) Box() {
	//TODO implement me
	panic("implement me")
}

func (p *brazilianCheesePizza) GetName() string {
	return p.Name
}

func (p *brazilianCheesePizza) GetDescription() string {
	return p.Description
}
