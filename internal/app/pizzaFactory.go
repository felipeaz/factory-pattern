package app

type PizzaFactory interface {
	CreatePizza(pizza string) *Pizza
}
