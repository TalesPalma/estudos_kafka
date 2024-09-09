package main

import kafkaservices "github.com/TalesPalma/kafka_consume/kafkaServices"

const (
	groupId = "validationGroup"
)

func main() {
	kafkaConsumer := kafkaservices.NewKafkaConsumer(groupId)
	kafkaConsumer.GetMessages()
}
