package app

type Pizza interface {
	Prepare() chan bool
	Cut()
	Box()
	GetName() string
	GetDescription() string
}
