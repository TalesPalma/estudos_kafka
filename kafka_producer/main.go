package main

import (
	"github.com/TalesPalma/messaging"
	"github.com/TalesPalma/models"
)

func main() {

	novaPessoa := models.Person{
		Name: "Tales",
		Age:  25,
	}

	msg, _ := novaPessoa.MarshalJson()

	producerMsg := messaging.NewKafkaProducer()
	producerMsg.SendMsg(msg, "my_topic")
	producerMsg.SendMsg([]byte("validar email"), "validation")
}
