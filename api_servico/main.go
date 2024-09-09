package main

import (
	kafkaservices "github.com/TalesPalma/kafkaServices"
	"github.com/TalesPalma/models"
)

func main() {

}

func SendMsgToKafka() {
	novaPessoa := models.Person{
		Name: "Tales",
		Age:  25,
	}

	msg, _ := novaPessoa.MarshalJson()

	producerMsg := kafkaservices.NewKafkaProducer()
	producerMsg.SendMsg(msg, "my_topic")

}
