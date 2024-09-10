package main

import (
	"github.com/TalesPalma/kafka_consume/database"
	kafkaservices "github.com/TalesPalma/kafka_consume/kafkaServices"
)

const (
	groupId = "validationGroup"
)

func main() {
	database.Connect()
	initkafka()
}

func initkafka() {
	kafkaConsumer := kafkaservices.NewOrderProcessor(groupId)
	kafkaConsumer.GetMessages()
}
