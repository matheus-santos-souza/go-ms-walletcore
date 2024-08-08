package handler

import (
	"fmt"
	"sync"

	"github.com/matheus-santos-souza/go-ms-walletcore/pkg/events"
	"github.com/matheus-santos-souza/go-ms-walletcore/pkg/kafka"
)

type BalanceUpdatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewBalanceUpdatedKafkaHandler(kafka *kafka.Producer) *BalanceUpdatedKafkaHandler {
	return &BalanceUpdatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *BalanceUpdatedKafkaHandler) Handle(message events.IEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(message, nil, "balances")
	fmt.Println("BalanceUpdatedKafkaHandler called")

}
