package app

type OrderManager interface {
	CreateOrder(pizza string) (order Order, err error)
	GetNextOrderInQueue() (pizza Pizza, err error)
}
