package kafkaservices

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	BootstrapServers = "localhost:29092,localhost:39092,localhost:49092" // Endereços dos brokers
	Topic            = "my_topic"                                        // Nome do topico
	RelatorioTopic   = "relatorio_topic"
	GerarRelatorio   = "Gerar Relatorio"
)

type OrderDispatcher struct {
	p     *kafka.Producer
	topic string
}

var Producer OrderDispatcher

func InitKafka() {
	Producer = NewOrderDispatcher()
}

func NewOrderDispatcher() OrderDispatcher {
	hostName, err := os.Hostname() // pega o nome do host para o kafka

	if err != nil {
		log.Fatalf("Não achou o host name")
	}

	// Cria o producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  BootstrapServers, // Endereços dos brokers
		"client.id":          hostName,         // nome do host como id para o kafka
		"acks":               "all",
		"retries":            5,
		"enable.auto.commit": true,
	})

	if err != nil {
		log.Fatalf("Erro ao criar o producer: %v", err)
	}

	// Retorna o producer
	return OrderDispatcher{
		p: p,
	}

}

func (producer *OrderDispatcher) SendMsg(msg []byte, gerarRelatorio bool) {
	topico := Topic
	err := producer.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topico, Partition: kafka.PartitionAny},
		Value:          msg},
		nil,
	)

	if err != nil {
		log.Fatalf("Erro ao enviar a mensagem para o primeiro topico: %v", err)
	}

	if gerarRelatorio {
		// Mandar msg parta o relatorio topic tambem
		topico = RelatorioTopic
		err = producer.p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topico, Partition: kafka.PartitionAny},
			Value:          msg},
			nil,
		)

		if err != nil {
			log.Fatalf("Erro ao enviar a mensagem para o relatorio topic: %v", err)
		}

	}

	// ele fica aguardando a mensagem ser enviada para o kafka e depois vai para a proxima
	go producer.handleEvents()

	//  limpa o buffer de mensagens do kafka para poder enviar novas mensagens
	producer.p.Flush(15 * 1000)
}

func (producer *OrderDispatcher) handleEvents() {
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
