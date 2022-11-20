package app

type OrderManager interface {
	AddToQueue(pizza Pizza) (positionInQueue int, err error)
	GetNextInQueue() (pizza Pizza, err error)
}
