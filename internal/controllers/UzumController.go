package controllers

import (
	"github.com/gin-gonic/gin"
	"gofax-billing/pkg/octo"
	"gofax-billing/pkg/utils"
	"net/http"
)

func UzumShopApiNotify(c *gin.Context) {
	var form octo.OctoNotifyResponse

	if err := c.ShouldBindJSON(&form); err != nil {
		utils.RespondJson(c, nil, http.StatusBadRequest, err.Error())
		return
	}

}
