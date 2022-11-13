package factory

import (
	"factory-pattern/internal/app"
	pizzaStore "factory-pattern/internal/domain/pizza"
	"factory-pattern/internal/domain/pizza/errors"
)

type brazilianPizzaFactory struct{}

func NewBrazilianPizzaFactory() app.PizzaFactory {
	return &brazilianPizzaFactory{}
}

func (b *brazilianPizzaFactory) CreatePizza(pizza string) (app.Pizza, error) {
	switch pizza {
	case pizzaStore.BrazilianCheesePizzaTitle:
		return pizzaStore.NewBrazilianCheesePizza(), nil
	case pizzaStore.BrazilianPepperoniPizzaTitle:
		return pizzaStore.NewBrazilianPepperoniPizza(), nil
	default:
		return nil, errors.NotFound(pizza)
	}
}
