package kafkaservices

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	BootstrapServers = "localhost:9092" // Endereços dos brokers
	Topic            = "my_topic"       // Nome do topico
)

type kafkaProducer struct {
	p     *kafka.Producer
	topic string
}

var Producer kafkaProducer

func InitKafka() {
	Producer = NewKafkaProducer()
}

func NewKafkaProducer() kafkaProducer {
	hostName, err := os.Hostname() // pega o nome do host para o kafka

	if err != nil {
		log.Fatalf("Não achou o host name")
	}

	// Cria o producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": BootstrapServers, // Endereços dos brokers
		"client.id":         hostName,         // nome do host como id para o kafka
	})

	if err != nil {
		log.Fatalf("Erro ao criar o producer: %v", err)
	}

	// Retorna o producer
	return kafkaProducer{
		p: p,
	}

}

func (producer *kafkaProducer) SendMsg(msg []byte) {

	topico := Topic
	err := producer.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topico, Partition: kafka.PartitionAny},
		Value:          msg},
		nil,
	)

	if err != nil {
		log.Fatalf("Erro ao enviar a mensagem: %v", err)
	}

	// ele fica aguardando a mensagem ser enviada para o kafka e depois vai para a proxima
	go producer.handleEvents()

	//  limpa o buffer de mensagens do kafka para poder enviar novas mensagens
	producer.p.Flush(15 * 1000)
}

func (producer *kafkaProducer) handleEvents() {
	// ele fica aguardando a mensagem ser enviada para o kafka e depois vai para a proxima

	// Pega todos os eventos
	for e := range producer.p.Events() {

		// paga o tipo de evento e faz algo
		switch ev := e.(type) {

		// caso tenha um evento de envio de mensagens
		case *kafka.Message:
			// verifica se deu tudo certo
			if ev.TopicPartition.Error != nil {
				log.Printf("Erro ao enviar mensagem: %v", ev.TopicPartition.Error)
			} else {
				log.Printf("Mensagem enviada: %s", ev.TopicPartition)
			}
		}

	}
}
