package kafkaservices

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	EstoqueTopic = "estoque_topic"
)

type KafkaProducer struct {
	p   *kafka.Producer
	err error
}

func NewkafkaProducer() KafkaProducer {
	p := KafkaProducer{}
	p.p, p.err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "estoque_producer",
	})

	return p
}

func (kprod *KafkaProducer) SendMessage(msg []byte) {

	go kprod.handlerEvents()

	estoqueTopic := EstoqueTopic
	err := kprod.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &estoqueTopic, Partition: kafka.PartitionAny},
		Value:          msg,
	},
		nil,
	)

	if err != nil {
		log.Println("Error producing message:", err)
	}

	kprod.p.Flush(15 * 1000)
}

func (kprod *KafkaProducer) handlerEvents() {
	for e := range kprod.p.Events() {
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
