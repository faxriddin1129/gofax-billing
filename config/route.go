package config

import (
	"github.com/gin-gonic/gin"
	payme "microservice/modules/payme/routes"
	"microservice/routes"
)

func RegisterRoutes(r *gin.Engine) {
	routes.MainRoutes(r)
	payme.PaymeRoutes(r)
}
