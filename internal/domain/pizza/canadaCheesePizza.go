package pizza

import "factory-pattern/internal/app"

const (
	CanadianCheesePizzaTitle       = "Cheese"
	CanadianCheesePizzaDescription = "Delicious Canadian Style Cheese Pizza"
)

type canadianCheesePizza struct {
	Name        string
	Description string
}

func NewCanadianCheesePizza() app.Pizza {
	return &canadianCheesePizza{
		Name:        CanadianCheesePizzaTitle,
		Description: CanadianCheesePizzaDescription,
	}
}

func (b *canadianCheesePizza) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (b *canadianCheesePizza) Cut() {
	//TODO implement me
	panic("implement me")
}

func (b *canadianCheesePizza) Box() {
	//TODO implement me
	panic("implement me")
}

func (b *canadianCheesePizza) IsPrepared() bool {
	//TODO implement me
	panic("implement me")
}
