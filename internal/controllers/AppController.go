package controllers

import (
	"gofax-billing/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	utils.RespondJson(c, nil, 200, "Welcome to the Asialuxe Billing")
}
