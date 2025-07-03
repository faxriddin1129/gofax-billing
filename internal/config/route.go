package config

import (
	"github.com/gin-gonic/gin"
	"microservice/routes"
)

func RegisterRoutes(r *gin.Engine) {
	routes.MainRoutes(r)
}
