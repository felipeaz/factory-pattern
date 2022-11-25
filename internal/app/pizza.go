package app

type Pizza interface {
	Prepare() (done bool)
	Cut()
	Box()
	GetName() string
	GetDescription() string
}
