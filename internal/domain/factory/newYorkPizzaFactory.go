package factory

import (
	"factory-pattern/internal/app"
	pizzaStore "factory-pattern/internal/domain/pizza"
	"factory-pattern/internal/domain/pizza/errors"
)

type newYorkPizzaFactory struct{}

func NewNYPizzaFactory() app.PizzaFactory {
	return &newYorkPizzaFactory{}
}

func (b *newYorkPizzaFactory) CreatePizza(pizza string) (app.Pizza, error) {
	switch pizza {
	case pizzaStore.NYCheesePizzaTitle:
		return pizzaStore.NewNYCheesePizza(), nil
	case pizzaStore.NYPepperoniPizzaTitle:
		return pizzaStore.NewNYPepperoniPizza(), nil
	default:
		return nil, errors.NotFound(pizza)
	}
}
