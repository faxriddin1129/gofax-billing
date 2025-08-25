package controllers

import (
	"gofax-billing/internal/requests"

	"github.com/gin-gonic/gin"
)

func PayByConfirmation(c *gin.Context) {
	requests.ConfirmationByCardValidate(c)
}

func FastPayGetLink(c *gin.Context) {
	requests.FastPayValidate(c)
}

func FastPayByCardGetLink(c *gin.Context) {
	requests.FastPayByCardValidate(c)
}
