package pizza

import (
	"fmt"
	"log"
	"time"

	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/pizza/constants"
)

const (
	BrazilianCheesePizzaTitle       = "Cheese"
	BrazilianCheesePizzaDescription = "Delicious Brazilian Style Cheese Pizza"
)

type brazilianCheesePizza struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewBrazilianCheesePizza() app.Pizza {
	return &brazilianCheesePizza{
		Name:               BrazilianCheesePizzaTitle,
		Description:        BrazilianCheesePizzaDescription,
		TimeToPrepareInSec: time.Second * 4,
	}
}

func (p *brazilianCheesePizza) Prepare(pizzaIsPrepared chan bool) {
	prepareTime := time.After(p.TimeToPrepareInSec)
	select {
	case <-prepareTime:
		pizzaIsPrepared <- true
		return
	default:
		log.Println(fmt.Sprintf(constants.PreparingPizzaText, p.Description))
	}
}

func (p *brazilianCheesePizza) Cut() {
	log.Println(fmt.Sprintf(constants.SlicingPizza, p.Name))
}

func (p *brazilianCheesePizza) Box() {
	log.Println(fmt.Sprintf(constants.BoxingPizza, p.Name))
}

func (p *brazilianCheesePizza) GetName() string {
	return p.Name
}

func (p *brazilianCheesePizza) GetDescription() string {
	return p.Description
}
