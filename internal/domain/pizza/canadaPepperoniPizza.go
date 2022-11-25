package pizza

import (
	"time"

	"factory-pattern/internal/app"
)

const (
	CanadianPepperoniPizzaTitle       = "Pepperoni"
	CanadianPepperoniPizzaDescription = "Delicious Canadian Style Pepperoni Pizza"
)

type canadianPepperoni struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewCanadianPepperoniPizza() app.Pizza {
	return &canadianPepperoni{
		Name:               CanadianPepperoniPizzaTitle,
		Description:        CanadianPepperoniPizzaDescription,
		TimeToPrepareInSec: time.Second * 3,
	}
}

func (p *canadianPepperoni) Prepare() bool {
	//TODO implement me
	panic("implement me")
}

func (p *canadianPepperoni) Cut() {
	//TODO implement me
	panic("implement me")
}

func (p *canadianPepperoni) Box() {
	//TODO implement me
	panic("implement me")
}

func (p *canadianPepperoni) GetName() string {
	return p.Name
}

func (p *canadianPepperoni) GetDescription() string {
	return p.Description
}
