package controllers

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/models"
	"microservice/pkg/utils"
	"net/http"
)

func Transactions(c *gin.Context) {
	transactions := models.TransactionGetAll()
	utils.RespondJson(c, transactions, http.StatusOK, "Ok")
	return
}
