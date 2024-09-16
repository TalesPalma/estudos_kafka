package main

import (
	"github.com/TalesPalma/controllers"
	"github.com/TalesPalma/database"
	kafkaservices "github.com/TalesPalma/kafkaServices"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	kafkaservices.InitKafka()
	InitGinServer()
}

func InitGinServer() {
	r := gin.Default()
	controllers.Headers(r)
	r.Run()
}
