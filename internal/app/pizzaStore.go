package app

type PizzaStore interface {
	Order(pizza string) (order Order, err error)
	Prepare() (pizza Pizza, err error)
}
