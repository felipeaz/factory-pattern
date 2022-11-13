package pizza

import "factory-pattern/internal/app"

const (
	BrazilianCheesePizzaTitle       = "Cheese"
	BrazilianCheesePizzaDescription = "Delicious Brazilian Style Cheese Pizza"
)

type brazilianCheesePizza struct {
	Name        string
	Description string
}

func NewBrazilianCheesePizza() app.Pizza {
	return &brazilianCheesePizza{
		Name:        BrazilianCheesePizzaTitle,
		Description: BrazilianCheesePizzaDescription,
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

func (b *brazilianCheesePizza) IsPrepared() bool {
	//TODO implement me
	panic("implement me")
}
