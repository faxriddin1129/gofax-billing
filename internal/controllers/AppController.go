package controllers

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/requests"
	"microservice/pkg/utils"
)

func Welcome(c *gin.Context) {
	utils.RespondJson(c, nil, 200, "Welcome to the Asialuxe Billing")
}

func FastPay(c *gin.Context) {
	requests.FastPayValidate(c)
}
