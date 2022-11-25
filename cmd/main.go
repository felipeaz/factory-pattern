package main

import "factory-pattern/internal/domain/pizzaStore"

func main() {
	brazilianPizzaStore := pizzaStore.NewBrazilianPizzaStore()
	brazilianPizzaStore.StartWork()
}
