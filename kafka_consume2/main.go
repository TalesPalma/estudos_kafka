package main

import (
	kafkaservices "github.com/TalesPalma/kafka_consume/kafka_services"
)

const (
	groupId = "myGroup2"
)

func main() {
	kafkaConsumer := kafkaservices.NewKafkaConsumer(groupId)
	kafkaConsumer.GetMessages()
}
