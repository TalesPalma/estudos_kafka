package kafkaservices

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	EstoqueTopic = "estoque_topic"
)

type OrderDispatcher struct {
	p   *kafka.Producer
	err error
}

func NewOrderDispatcher() OrderDispatcher {
	p := OrderDispatcher{}
	p.p, p.err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "estoque_producer",
	})

	return p
}

func (order *OrderDispatcher) SendMessage(msg []byte) {

	go order.handlerEvents()

	estoqueTopic := EstoqueTopic
	err := order.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &estoqueTopic, Partition: kafka.PartitionAny},
		Value:          msg,
	},
		nil,
	)

	if err != nil {
		log.Println("Error producing message:", err)
	}

	order.p.Flush(15 * 1000)
}

func (order *OrderDispatcher) handlerEvents() {
	for e := range order.p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Failed to deliver message: ", ev.TopicPartition.Error)
			} else {
				fmt.Println("Message delired to topic: ", ev.TopicPartition.Topic)
			}
		}
	}
}
