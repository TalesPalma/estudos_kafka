package main

import kafkaservices "www.github.com/TalesPalma/kafkaServices"

func main() {
	order := kafkaservices.NewOrderProcessor("order-processor")
	order.GetMessages()
}
