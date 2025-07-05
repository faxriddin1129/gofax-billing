package controllers

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/requests"
)

func FastPayGetLink(c *gin.Context) {
	requests.FastPayValidate(c)
}
