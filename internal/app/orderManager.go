package app

type OrderManager interface {
	CreateOrder(pizza string) (err error)
	GetNextOrderInQueue() (pizza Pizza, err error)
}
