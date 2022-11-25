package pizza

import (
	"time"

	"factory-pattern/internal/app"
)

const (
	NYPepperoniPizzaTitle       = "Pepperoni"
	NYPepperoniPizzaDescription = "Delicious New York Style Pepperoni Pizza"
)

type newYorkPepperoniPizza struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewNYPepperoniPizza() app.Pizza {
	return &newYorkPepperoniPizza{
		Name:               NYPepperoniPizzaTitle,
		Description:        NYPepperoniPizzaDescription,
		TimeToPrepareInSec: time.Second * 8,
	}
}

func (p *newYorkPepperoniPizza) Prepare() bool {
	//TODO implement me
	panic("implement me")
}

func (p *newYorkPepperoniPizza) Cut() {
	//TODO implement me
	panic("implement me")
}

func (p *newYorkPepperoniPizza) Box() {
	//TODO implement me
	panic("implement me")
}

func (p *newYorkPepperoniPizza) GetName() string {
	return p.Name
}

func (p *newYorkPepperoniPizza) GetDescription() string {
	return p.Description
}
