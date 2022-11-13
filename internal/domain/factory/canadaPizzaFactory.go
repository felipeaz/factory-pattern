package factory

import (
	"factory-pattern/internal/app"
	pizzaStore "factory-pattern/internal/domain/pizza"
	"factory-pattern/internal/domain/pizza/errors"
)

type canadianPizzaFactory struct{}

func NewCanadianPizzaFactory() app.PizzaFactory {
	return &canadianPizzaFactory{}
}

func (b *canadianPizzaFactory) CreatePizza(pizza string) (app.Pizza, error) {
	switch pizza {
	case pizzaStore.CanadianCheesePizzaTitle:
		return pizzaStore.NewCanadianCheesePizza(), nil
	case pizzaStore.CanadianPepperoniPizzaTitle:
		return pizzaStore.NewCanadianPepperoniPizza(), nil
	default:
		return nil, errors.NotFound(pizza)
	}
}
