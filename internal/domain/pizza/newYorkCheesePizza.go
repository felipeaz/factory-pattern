package pizza

import (
	"fmt"
	"log"
	"time"

	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/pizza/constants"
)

const (
	NYCheesePizzaTitle       = "Cheese"
	NYCheesePizzaDescription = "Delicious New York Style Cheese Pizza"
)

type newYorkCheesePizza struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewNYCheesePizza() app.Pizza {
	return &newYorkCheesePizza{
		Name:               NYCheesePizzaTitle,
		Description:        NYCheesePizzaDescription,
		TimeToPrepareInSec: time.Second * 6,
	}
}

func (p *newYorkCheesePizza) Prepare(pizzaIsPrepared chan bool) {
	prepareTime := time.After(p.TimeToPrepareInSec)
	select {
	case <-prepareTime:
		pizzaIsPrepared <- true
		return
	default:
		log.Println(fmt.Sprintf(constants.PreparingPizzaText, p.Description))
	}
}

func (p *newYorkCheesePizza) Cut() {
	log.Println(fmt.Sprintf(constants.SlicingPizza, p.Name))
}

func (p *newYorkCheesePizza) Box() {
	log.Println(fmt.Sprintf(constants.BoxingPizza, p.Name))
}

func (p *newYorkCheesePizza) GetName() string {
	return p.Name
}

func (p *newYorkCheesePizza) GetDescription() string {
	return p.Description
}
