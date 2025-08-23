package controllers

import (
	"github.com/gin-gonic/gin"
	"gofax-billing/internal/requests"
)

func PayByConfirmation(c *gin.Context) {
	requests.ConfirmationByCardValidate(c)
}
