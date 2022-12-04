package pizza

import (
	"fmt"
	"log"
	"time"

	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/pizza/constants"
)

const (
	CanadianPepperoniPizzaTitle       = "Pepperoni"
	CanadianPepperoniPizzaDescription = "Delicious Canadian Style Pepperoni Pizza"
)

type canadianPepperoni struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewCanadianPepperoniPizza() app.Pizza {
	return &canadianPepperoni{
		Name:               CanadianPepperoniPizzaTitle,
		Description:        CanadianPepperoniPizzaDescription,
		TimeToPrepareInSec: time.Second * 3,
	}
}

func (p *canadianPepperoni) Prepare(pizzaIsPrepared chan bool) {
	prepareTime := time.After(p.TimeToPrepareInSec)
	select {
	case <-prepareTime:
		pizzaIsPrepared <- true
		return
	default:
		log.Println(fmt.Sprintf(constants.PreparingPizzaText, p.Description))
	}
}

func (p *canadianPepperoni) Cut() {
	log.Println(fmt.Sprintf(constants.SlicingPizza, p.Name))
}

func (p *canadianPepperoni) Box() {
	log.Println(fmt.Sprintf(constants.BoxingPizza, p.Name))
}

func (p *canadianPepperoni) GetName() string {
	return p.Name
}

func (p *canadianPepperoni) GetDescription() string {
	return p.Description
}
