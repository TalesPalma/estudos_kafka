package kafkaservices

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	bootstrapServers = "localhost:9092"
	topic            = "my_topic"
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
	if groupId == "myGroup1" {
		err = c.SubscribeTopics([]string{topic}, balaceAdorConsumer)

	} else {
		err = c.SubscribeTopics([]string{"validation"}, nil)
	}

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
		// else if check for timeout erro
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
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
