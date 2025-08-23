package controllers

import (
	"github.com/gin-gonic/gin"
	"gofax-billing/pkg/utils"
	"net/http"
)

func PayByConfirmation(c *gin.Context) {
	utils.RespondJson(c, nil, http.StatusOK, "OK")
}
