package pizza

import (
	"time"

	"factory-pattern/internal/app"
)

const (
	NYPepperoniPizzaTitle       = "Pepperoni"
	NYPepperoniPizzaDescription = "Delicious New York Style Pepperoni Pizza"
)

type newYorkPepperoniPizza struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewNYPepperoniPizza() app.Pizza {
	return &newYorkPepperoniPizza{
		Name:               NYPepperoniPizzaTitle,
		Description:        NYPepperoniPizzaDescription,
		TimeToPrepareInSec: time.Second * 8,
	}
}

func (b *newYorkPepperoniPizza) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (b *newYorkPepperoniPizza) Cut() {
	//TODO implement me
	panic("implement me")
}

func (b *newYorkPepperoniPizza) Box() {
	//TODO implement me
	panic("implement me")
}

func (b *newYorkPepperoniPizza) IsPrepared() bool {
	//TODO implement me
	panic("implement me")
}
