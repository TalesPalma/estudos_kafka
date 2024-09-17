package kafkaservices

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
)

const (
	EstoqueTopic = "estoque_topic"
)

type OrderDispatcher struct {
	p   *kafka.Producer
	err error
}

func NewOrderDispatcher() OrderDispatcher {
	order := OrderDispatcher{}
	order.p, order.err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,   // Endereços dos brokers
		"client.id":         "estoque_producer", // ID do consumidor
		"acks":              "all",              // Esperar todas as replicas do cluster está em syncronia
	})

	return order
}

func (order *OrderDispatcher) SendMessage(msg []byte) {

	go order.handlerEvents()

	estoqueTopic := EstoqueTopic

	correlationId := generateCorrelationId()
	key := correlationId

	//Criar mensagem com correlationId
	msgKafka := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &estoqueTopic, Partition: kafka.PartitionAny},
		Value:          msg,
		Key:            []byte(key),
		Headers:        []kafka.Header{{Key: "mycorrelationid", Value: []byte(correlationId)}},
	}

	//Enviar mensagem para o kafka
	err := order.p.Produce(&msgKafka, nil)

	//Tratamento de erro caso a mensagem não seja enviada
	if err != nil {
		log.Println("Error producing message:", err)
	}

	// Flush messages depois de 15 segundos
	flushTimeOut := 15 * time.Second
	if remaining := order.p.Flush(int(flushTimeOut)); remaining > 0 {
		log.Println("Remaining:", remaining)
	}

	// Fechar producer
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

// Gerador de correlation id
func generateCorrelationId() string {
	return uuid.New().String()
}
