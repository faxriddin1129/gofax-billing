package controllers

import (
	"gofax-billing/pkg/click"
	"gofax-billing/pkg/octo"
	"gofax-billing/pkg/payme"
	"gofax-billing/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OctoShopApiNotify(c *gin.Context) {
	var form octo.OctoNotifyResponse

	if err := c.ShouldBindJSON(&form); err != nil {
		utils.RespondJson(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	octo.NotifyShopApi(&form, c)
}

func PaymeShopApiNotify(c *gin.Context) {
	var form payme.PaymeRequest

	if err := c.ShouldBindJSON(&form); err != nil {
		utils.RespondJson(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	payme.NotifyShopApi(&form, c)
}

func ClickShopApiNotify(c *gin.Context) {
	var form click.ResponseShopApi

	if err := c.ShouldBind(&form); err != nil {
		utils.RespondJson(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	click.NotifyShopApi(&form, c)
}
