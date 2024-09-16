package kafkaservices

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"www.github.com/TalesPalma/relatorio"
)

const (
	bootstrapServers = "localhost:29092,localhost:39092,localhost:49092" // Endereços dos brokers
	RelatorioTopic   = "relatorio_topic"
)

type OrderProcessor struct {
	consumer *kafka.Consumer
}

func NewOrderProcessor(groupId string) *OrderProcessor {
	//Create consumer
	order, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
		// Configurações SSL
		"security.protocol": "PLAINTEXT",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	//Subscribe to topic my_topic and estoque_topic
	err = order.SubscribeTopics([]string{RelatorioTopic}, nil)

	if err != nil {
		log.Fatalf("Failed to subscribe to topic my topic and estoque_topic: %s", err)
	}

	log.Println("Registred on topics my topic and estoque_topic ...")

	return &OrderProcessor{
		consumer: order,
	}
}

func (order *OrderProcessor) GetMessages() {
	run := true
	for run {
		msg, err := order.consumer.ReadMessage(time.Second * 3) // wait for message for 1 second

		if err == nil {
			processMessage(msg)
			continue
		}

	}

	//Close consumer
	err := order.consumer.Close()
	if err != nil {
		log.Fatalf("Failed to close consumer: %s", err)
	}

}

// Balacea o consumer nas particoes atualizadas pelo kafka
func balaceAdorConsumer(order *kafka.Consumer, e kafka.Event) error {
	switch ev := e.(type) {
	case kafka.AssignedPartitions:
		fmt.Printf("Assigned partitions: %v\n", ev.Partitions)
		order.Assign(ev.Partitions)
	case kafka.RevokedPartitions:
		fmt.Printf("Revoked partitions: %v\n", ev.Partitions)
		order.Unassign()
	}
	return nil
}

func processMessage(msg *kafka.Message) {
	log.Printf("Log message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	relatorio.GenerateRelatorio(msg.Value)
}
