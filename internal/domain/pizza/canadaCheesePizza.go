package pizza

import (
	"fmt"
	"log"
	"time"

	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/pizza/constants"
)

const (
	CanadianCheesePizzaTitle       = "Cheese"
	CanadianCheesePizzaDescription = "Delicious Canadian Style Cheese Pizza"
)

type canadianCheesePizza struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewCanadianCheesePizza() app.Pizza {
	return &canadianCheesePizza{
		Name:               CanadianCheesePizzaTitle,
		Description:        CanadianCheesePizzaDescription,
		TimeToPrepareInSec: time.Second * 2,
	}
}

func (p *canadianCheesePizza) Prepare(pizzaIsPrepared chan bool) {
	prepareTime := time.After(p.TimeToPrepareInSec)
	select {
	case <-prepareTime:
		pizzaIsPrepared <- true
		return
	default:
		log.Println(fmt.Sprintf(constants.PreparingPizzaText, p.Description))
	}
}

func (p *canadianCheesePizza) Cut() {
	log.Println(fmt.Sprintf(constants.SlicingPizza, p.Name))
}

func (p *canadianCheesePizza) Box() {
	log.Println(fmt.Sprintf(constants.BoxingPizza, p.Name))
}

func (p *canadianCheesePizza) GetName() string {
	return p.Name
}

func (p *canadianCheesePizza) GetDescription() string {
	return p.Description
}
