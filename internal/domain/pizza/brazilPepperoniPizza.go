package pizza

import (
	"fmt"
	"log"
	"time"

	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/pizza/constants"
)

const (
	BrazilianPepperoniPizzaTitle       = "Pepperoni"
	BrazilianPepperoniPizzaDescription = "Delicious Brazilian Style Pepperoni Pizza"
)

type brazilianPepperoni struct {
	Name               string
	Description        string
	TimeToPrepareInSec time.Duration
}

func NewBrazilianPepperoniPizza() app.Pizza {
	return &brazilianPepperoni{
		Name:               BrazilianPepperoniPizzaTitle,
		Description:        BrazilianPepperoniPizzaDescription,
		TimeToPrepareInSec: time.Second * 5,
	}
}

func (p *brazilianPepperoni) Prepare(pizzaIsPrepared chan bool) {
	prepareTime := time.After(p.TimeToPrepareInSec)
	select {
	case <-prepareTime:
		pizzaIsPrepared <- true
		return
	default:
		log.Println(fmt.Sprintf(constants.PreparingPizzaText, p.Description))
	}
}

func (p *brazilianPepperoni) Cut() {
	log.Println(fmt.Sprintf(constants.SlicingPizza, p.Name))
}

func (p *brazilianPepperoni) Box() {
	log.Println(fmt.Sprintf(constants.BoxingPizza, p.Name))
}

func (p *brazilianPepperoni) GetName() string {
	return p.Name
}

func (p *brazilianPepperoni) GetDescription() string {
	return p.Description
}
