package app

import "time"

type Pizza interface {
	Prepare()
	Cut()
	Box()
	GetTimeToPrepareInSeconds() time.Duration
	IsPrepared() bool
}
