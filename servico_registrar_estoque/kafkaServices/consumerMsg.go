package kafkaservices

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/TalesPalma/kafka_consume/database"
	"github.com/TalesPalma/kafka_consume/database/models"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gorm.io/gorm"
)

const (
	bootstrapServers = "localhost:9092"
	Topic            = "my_topic"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(groupId string) *KafkaConsumer {
	//Create consumer
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	//Subscribe to topic
	err = c.SubscribeTopics([]string{Topic}, nil)

	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	return &KafkaConsumer{
		consumer: c,
	}
}

func (k *KafkaConsumer) GetMessages() {
	run := true
	for run {
		msg, err := k.consumer.ReadMessage(time.Second) // wait for message for 1 second

		// check for errors and print message
		// else if check for timeout err
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			insertProductDatabase(msg.Value)
		} else if !err.(kafka.Error).IsTimeout() {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	//Close consumer
	err := k.consumer.Close()
	if err != nil {
		log.Fatalf("Failed to close consumer: %s", err)
	}

}

// Balacea o consumer nas particoes atualizadas pelo kafka
func balaceAdorConsumer(c *kafka.Consumer, e kafka.Event) error {
	switch ev := e.(type) {
	case kafka.AssignedPartitions:
		fmt.Printf("Assigned partitions: %v\n", ev.Partitions)
		c.Assign(ev.Partitions)
	case kafka.RevokedPartitions:
		fmt.Printf("Revoked partitions: %v\n", ev.Partitions)
		c.Unassign()
	}
	return nil
}

// Funcao para inserir o novo produto ou atualizar o estoque
// Ela deserializa o json e insere ou atualiza o estoque
func insertProductDatabase(msg []byte) {
	var product models.Product
	product.Unmarshal(msg)
	verificarSeOProdutoJaEstaRegistradoNoBanco(product)
}

// Se o produto já existe então será atualizado o estoque e se não existe ele será criado
// E caso ele existe, ele será somando ao estoque o valor da quantidade adicionada
func verificarSeOProdutoJaEstaRegistradoNoBanco(product models.Product) {
	kafkaProducer := NewkafkaProducer()
	productDB, err := buscarProdutoNoBanco(product.Name)

	if err != nil {
		log.Println("Error when trying to find product", err)
		return
	}

	if productDB != nil {
		product.Stock = product.Stock + product.Qty
		database.DB.Save(&product)
		log.Println("Product updated:", product)
		kafkaProducer.SendMessage([]byte("Product updated: " + product.Name))

	} else {
		database.DB.Create(&product)
		log.Println("Product created: ", product)
		kafkaProducer.SendMessage([]byte("New Product created: " + product.Name))
	}
}

func buscarProdutoNoBanco(s string) (*models.Product, error) {
	var productDB models.Product
	result := database.DB.Where("name = ?", s).First(&productDB)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Produto não encontrado ou seja sera criado", result.Error)
			return nil, nil
		}

		log.Println("Product not found: ", productDB)
		return nil, result.Error
	}
	return &productDB, nil
}
