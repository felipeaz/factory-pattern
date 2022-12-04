package app

type Pizza interface {
	Prepare(pizzaIsPrepared chan bool)
	Cut()
	Box()
	GetName() string
	GetDescription() string
}
