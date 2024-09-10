package main

import kafkaservices "github.com/TalesPalma/kafka_consume2/kafkaServices"

const (
	groupId = "logGroup"
)

func main() {
	kafkaConsumer := kafkaservices.NewOrderProcessor(groupId)
	kafkaConsumer.GetMessages()
}
