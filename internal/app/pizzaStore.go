package app

type PizzaStore interface {
	Order(pizza string) (err error)
	Prepare() (pizza Pizza, err error)
}
