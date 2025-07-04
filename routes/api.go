package routes

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/controllers"
)

func MainRoutes(r *gin.Engine) {
	r.GET("/", controllers.Welcome)
	r.GET("/transaction", controllers.Transactions)
	r.POST("/fast-pay", controllers.FastPay)
}
