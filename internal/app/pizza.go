package app

type Pizza interface {
	Prepare()
	Cut()
	Box()
	IsPrepared() bool
}
