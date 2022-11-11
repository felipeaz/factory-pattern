package app

type PizzaStore interface {
	Order(pizza string)
	Prepare(order Order) *Pizza
}
