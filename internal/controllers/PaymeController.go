package controllers

import (
	"github.com/gin-gonic/gin"
	"gofax-billing/pkg/payme"
	"gofax-billing/pkg/utils"
	"net/http"
)

func PaymeShopApiNotify(c *gin.Context) {
	var form payme.PaymeRequest

	if err := c.ShouldBindJSON(&form); err != nil {
		utils.RespondJson(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	payme.NotifyShopApi(&form, c)
}
