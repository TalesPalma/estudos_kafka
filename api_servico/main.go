package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/TalesPalma/controllers"
	"github.com/TalesPalma/database"
	"github.com/TalesPalma/database/models"
	kafkaservices "github.com/TalesPalma/kafkaServices"
	"github.com/gin-gonic/gin"
)

func main() {
	var wg sync.WaitGroup

	database.InitDatabase()
	kafkaservices.InitKafka()

	wg.Add(1)
	go InitGinServer(&wg)

	go gerarRelatorioPorDia()

	wg.Wait()
}

func gerarRelatorioPorDia() {
	for {
		var products []models.Product
		database.DB.Find(&products)
		allProducts, _ := json.Marshal(products)
		kafkaservices.Producer.SendMsg([]byte(allProducts), true)
		time.Sleep(24 * time.Hour)
		fmt.Println("Relatorio Gerado")
	}
}

func InitGinServer(wg *sync.WaitGroup) {
	r := gin.Default()
	controllers.Headers(r)
	r.Run()
	wg.Done()
}
