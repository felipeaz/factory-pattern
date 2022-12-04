package pizzaStore

import (
	"log"

	"factory-pattern/internal/app"
)

type worker struct {
	doneWorking chan bool
	pizzaQueue  chan string
}

func initializeWorker() worker {
	return worker{
		doneWorking: make(chan bool),
		pizzaQueue:  make(chan string),
	}
}

func (w worker) startWork(orderManager app.OrderManager, pizzaFactory app.PizzaFactory) {
	defer w.stopWork()

	go orderManager.GetNextOrderInQueue(w.pizzaQueue, w.doneWorking)

	for orderedPizza := range w.pizzaQueue {
		pizzaIsPrepared := make(chan bool)

		pizzaInProduction, err := pizzaFactory.CreatePizza(orderedPizza)
		if err != nil {
			w.doneWorking <- true
		}

		pizzaInProduction.Prepare(pizzaIsPrepared)
		select {
		case <-pizzaIsPrepared:
			log.Printf("%s Pizza is ready for delivery", pizzaInProduction.GetName())
			finishPizza(pizzaInProduction)
			continue
		default:
			log.Printf("Pizza will be ready to delivery soon")
		}
	}
}

func (w worker) stopWork() {
	close(w.doneWorking)
	close(w.pizzaQueue)
}

func finishPizza(pizza app.Pizza) {
	pizza.Cut()
	pizza.Box()
	delivery(pizza)
}

func delivery(pizza app.Pizza) {
	log.Printf("A %s Pizza is out for delivery", pizza.GetDescription())
}
