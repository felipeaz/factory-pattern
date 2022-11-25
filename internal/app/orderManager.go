package app

type OrderManager interface {
	CreateOrder(pizza string) (err error)
	GetNextOrderInQueue(queue chan<- string, done <-chan bool)
}
