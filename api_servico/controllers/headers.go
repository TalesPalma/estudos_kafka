package controllers

import (
	"github.com/gin-gonic/gin"
)

func Headers(r *gin.Engine) {
	r.GET("/", Index)
	r.GET("/products", GetAllProducts)
	r.POST("/products", PostProducts)
	r.DELETE("/products/:code", DeleteProductsByCode)
	r.PUT("/products/:code", PutProductsByCode)
}
