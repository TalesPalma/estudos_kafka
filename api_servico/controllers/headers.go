package controllers

import (
	"github.com/gin-gonic/gin"
)

func Headers(r *gin.Engine) {
	r.GET("/", Index)
}
