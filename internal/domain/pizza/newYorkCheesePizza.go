package pizza

import "factory-pattern/internal/app"

const (
	NYCheesePizzaTitle       = "Cheese"
	NYCheesePizzaDescription = "Delicious New York Style Cheese Pizza"
)

type newYorkCheesePizza struct {
	Name        string
	Description string
}

func NewNYCheesePizza() app.Pizza {
	return &newYorkCheesePizza{
		Name:        NYCheesePizzaTitle,
		Description: NYCheesePizzaDescription,
	}
}

func (b *newYorkCheesePizza) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (b *newYorkCheesePizza) Cut() {
	//TODO implement me
	panic("implement me")
}

func (b *newYorkCheesePizza) Box() {
	//TODO implement me
	panic("implement me")
}

func (b *newYorkCheesePizza) IsPrepared() bool {
	//TODO implement me
	panic("implement me")
}
