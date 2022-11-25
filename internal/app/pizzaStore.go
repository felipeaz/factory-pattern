package app

type PizzaStore interface {
	StartWork()
	Order(pizza string) (err error)
}
