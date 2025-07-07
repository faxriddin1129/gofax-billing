package controllers

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/requests"
)

func FastPayGetLink(c *gin.Context) {
	requests.FastPayValidate(c)
}

func FastPayByCardGetLink(c *gin.Context) {
	requests.FastPayByCardValidate(c)
}
