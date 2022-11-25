package order

import (
	"log"

	kafkalib "github.com/confluentinc/confluent-kafka-go/kafka"

	"factory-pattern/infrastructure/kafka"
	"factory-pattern/infrastructure/kafka/event"
	"factory-pattern/internal/app"
	"factory-pattern/internal/errors"
)

type orderManager struct {
	kafkaHandler kafka.Kafka
	topic        string
	groupId      string
}

func NewManager(kafkaHandler kafka.Kafka) app.OrderManager {
	return &orderManager{
		kafkaHandler: kafkaHandler,
	}
}

func (o *orderManager) CreateOrder(pizza string) (err error) {
	orderedPizza, err := NewOrderInBytes(pizza)
	if err != nil {
		return errors.WithStack("failed to create order in bytes", err)
	}

	err = o.placeOrderInQueue(orderedPizza)
	if err != nil {
		return errors.WithStack("failed to place an order in queue", err)
	}

	return err
}

func (o *orderManager) placeOrderInQueue(order []byte) error {
	kafkaEvent, err := o.kafkaHandler.Produce(order)
	if err != nil {
		return errors.WithStack("error calling producer", err)
	}

	msg, err := event.HandleEvent(kafkaEvent)
	if err != nil {
		return errors.WithStack("error producing message", err)
	}

	log.Println("Order placed at:", msg.Timestamp)

	return nil
}

func (o *orderManager) GetNextOrderInQueue(queue chan<- string, done <-chan bool) {
	orderMessageChan := make(chan *kafkalib.Message)
	defer close(orderMessageChan)

	go o.kafkaHandler.Consume(orderMessageChan, done)

	select {
	case <-done:
		log.Println("stopped GetNextOrderInQueue")
		return
	case msg := <-orderMessageChan:
		orderDto, err := NewOrderDtoFromBytes(msg.Value)
		if err != nil {
			return
		}
		queue <- orderDto.Pizza
	}
}
