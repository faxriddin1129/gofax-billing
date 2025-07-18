package routes

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/controllers"
)

func MainRoutes(r *gin.Engine) {

	r.GET("/api", controllers.Welcome)

	fp := r.Group("/api/fast-pay")
	{
		fp.POST("/get-link", controllers.FastPayGetLink)
		fp.POST("/get-link-by-card", controllers.FastPayByCardGetLink)
	}

	notify := r.Group("/api")
	{
		notify.POST("/octo/notify", controllers.OctoShopApiNotify)
		notify.POST("/payme/notify", controllers.PaymeShopApiNotify)
		notify.POST("/click/notify-prepare", controllers.PaymeShopApiNotify)
		notify.POST("/click/notify-complete", controllers.PaymeShopApiNotify)
		notify.POST("/uzum/notify", controllers.OctoShopApiNotify)
	}
}
