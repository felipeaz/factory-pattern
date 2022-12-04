package pizza

import (
	"fmt"
	"log"
	"time"

	"factory-pattern/internal/app"
	"factory-pattern/internal/domain/pizza/constants"
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

func (p *newYorkPepperoniPizza) Prepare(pizzaIsPrepared chan bool) {
	prepareTime := time.After(p.TimeToPrepareInSec)
	select {
	case <-prepareTime:
		pizzaIsPrepared <- true
		return
	default:
		log.Println(fmt.Sprintf(constants.PreparingPizzaText, p.Description))
	}
}

func (p *newYorkPepperoniPizza) Cut() {
	log.Println(fmt.Sprintf(constants.SlicingPizza, p.Name))
}

func (p *newYorkPepperoniPizza) Box() {
	log.Println(fmt.Sprintf(constants.BoxingPizza, p.Name))
}

func (p *newYorkPepperoniPizza) GetName() string {
	return p.Name
}

func (p *newYorkPepperoniPizza) GetDescription() string {
	return p.Description
}
