package pizza

import (
	"time"

	"factory-pattern/internal/app"
)

const (
	BrazilianPepperoniPizzaTitle       = "Pepperoni"
	BrazilianPepperoniPizzaDescription = "Delicious Brazilian Style Pepperoni Pizza"
)

type brazilianPepperoni struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewBrazilianPepperoniPizza() app.Pizza {
	return &brazilianPepperoni{
		Name:               BrazilianPepperoniPizzaTitle,
		Description:        BrazilianPepperoniPizzaDescription,
		TimeToPrepareInSec: time.Second * 5,
	}
}

func (p *brazilianPepperoni) Prepare() chan bool {
	//TODO implement me
	panic("implement me")
}

func (p *brazilianPepperoni) Cut() {
	//TODO implement me
	panic("implement me")
}

func (p *brazilianPepperoni) Box() {
	//TODO implement me
	panic("implement me")
}

func (p *brazilianPepperoni) GetName() string {
	return p.Name
}

func (p *brazilianPepperoni) GetDescription() string {
	return p.Description
}
