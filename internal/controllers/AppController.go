package controllers

import (
	"github.com/gin-gonic/gin"
	"microservice/pkg/utils"
)

func Welcome(c *gin.Context) {
	utils.RespondWithSuccess(c, nil, 200, "Welcome to the Asialuxe Billing")
}
