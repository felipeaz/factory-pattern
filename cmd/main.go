package main

import (
	"log"

	"factory-pattern/internal/domain/pizza"
	"factory-pattern/internal/domain/pizzaStore"
)

func main() {
	brazilianPizzaStore := pizzaStore.NewBrazilianPizzaStore()
	brazilianPizzaStore.StartWork()

	err := brazilianPizzaStore.Order(pizza.BrazilianCheesePizzaTitle)
	if err != nil {
		log.Fatal("failed to order pizza", err)
	}

	err = brazilianPizzaStore.Order(pizza.BrazilianPepperoniPizzaTitle)
	if err != nil {
		log.Fatal("failed to order pizza", err)
	}
}
