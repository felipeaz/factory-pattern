package pizza

import "factory-pattern/internal/app"

const (
	BrazilianPepperoniPizzaTitle       = "Pepperoni"
	BrazilianPepperoniPizzaDescription = "Delicious Brazilian Style Pepperoni Pizza"
)

type brazilianPepperoni struct {
	Name        string
	Description string
}

func NewBrazilianPepperoniPizza() app.Pizza {
	return &brazilianPepperoni{
		Name:        BrazilianPepperoniPizzaTitle,
		Description: BrazilianPepperoniPizzaDescription,
	}
}

func (b *brazilianPepperoni) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianPepperoni) Cut() {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianPepperoni) Box() {
	//TODO implement me
	panic("implement me")
}

func (b *brazilianPepperoni) IsPrepared() bool {
	//TODO implement me
	panic("implement me")
}
