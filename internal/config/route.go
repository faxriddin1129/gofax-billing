package config

import (
	"github.com/gin-gonic/gin"
	"gofax-billing/routes"
)

func RegisterRoutes(r *gin.Engine) {
	routes.MainRoutes(r)
}
