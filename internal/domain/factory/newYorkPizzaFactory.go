package factory

import "factory-pattern/internal/app"

type NewYorkPizzaFactory struct {
}

func (b *NewYorkPizzaFactory) CreatePizza(pizza string) *app.Pizza {
	//TODO implement me
	panic("implement me")
}
