package pizza

import "factory-pattern/internal/app"

const (
	CanadianPepperoniPizzaTitle       = "Pepperoni"
	CanadianPepperoniPizzaDescription = "Delicious Canadian Style Pepperoni Pizza"
)

type canadianPepperoni struct {
	Name        string
	Description string
}

func NewCanadianPepperoniPizza() app.Pizza {
	return &canadianPepperoni{
		Name:        CanadianPepperoniPizzaTitle,
		Description: CanadianPepperoniPizzaDescription,
	}
}

func (b *canadianPepperoni) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (b *canadianPepperoni) Cut() {
	//TODO implement me
	panic("implement me")
}

func (b *canadianPepperoni) Box() {
	//TODO implement me
	panic("implement me")
}

func (b *canadianPepperoni) IsPrepared() bool {
	//TODO implement me
	panic("implement me")
}
