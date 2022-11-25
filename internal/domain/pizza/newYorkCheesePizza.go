package pizza

import (
	"time"

	"factory-pattern/internal/app"
)

const (
	NYCheesePizzaTitle       = "Cheese"
	NYCheesePizzaDescription = "Delicious New York Style Cheese Pizza"
)

type newYorkCheesePizza struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewNYCheesePizza() app.Pizza {
	return &newYorkCheesePizza{
		Name:               NYCheesePizzaTitle,
		Description:        NYCheesePizzaDescription,
		TimeToPrepareInSec: time.Second * 6,
	}
}

func (p *newYorkCheesePizza) Prepare() chan bool {
	//TODO implement me
	panic("implement me")
}

func (p *newYorkCheesePizza) Cut() {
	//TODO implement me
	panic("implement me")
}

func (p *newYorkCheesePizza) Box() {
	//TODO implement me
	panic("implement me")
}

func (p *newYorkCheesePizza) GetName() string {
	return p.Name
}

func (p *newYorkCheesePizza) GetDescription() string {
	return p.Description
}
