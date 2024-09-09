package controllers

import (
	"net/http"

	"github.com/TalesPalma/database"
	"github.com/TalesPalma/database/models"
	kafkaservices "github.com/TalesPalma/kafkaServices"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}

func PostProducts(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)

	if product.IsEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing fields",
		})
		return
	}

	error := database.DB.Save(&product)
	if error.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error,
		})
	} else {
		c.JSON(http.StatusCreated, product)
		kafkaservices.Producer.SendMsg(product.MarshalJson())
	}

}

func DeleteProductsByCode(c *gin.Context) {
	var product models.Product
	code := c.Param("code")
	error := database.DB.Where("code = ?", code).Delete(&product)

	if error.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error,
		})
	} else {
		c.JSON(http.StatusOK, product)
	}

}

func PutProductsByCode(c *gin.Context) {
	var product models.Product

	code := c.Param("code")
	c.BindJSON(&product)

	error := database.DB.Where("code = ?", code).Updates(&product)

	if error.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error,
		})
	} else {
		c.JSON(http.StatusOK, product)
	}

}
