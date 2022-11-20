package app

type PizzaStore interface {
	Order(pizza string) (positionInQueue int, err error)
	Prepare(order OrderManager) (orderedPizza *Pizza, err error)
}
