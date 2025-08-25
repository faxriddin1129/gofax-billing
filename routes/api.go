package routes

import (
	"gofax-billing/internal/controllers"

	"github.com/gin-gonic/gin"
)

func MainRoutes(r *gin.Engine) {

	r.GET("/api", controllers.Welcome)

	r.POST("/api/fast-pay/get-link", controllers.FastPayGetLink)
	r.POST("/api/fast-pay/get-link-by-card", controllers.FastPayByCardGetLink)
	r.POST("/api/confirmation-pay", controllers.PayByConfirmation)

	r.POST("/api/octo/notify", controllers.OctoShopApiNotify)
	r.POST("/api/payme/notify", controllers.PaymeShopApiNotify)
	r.POST("/api/click/notify", controllers.ClickShopApiNotify)
}
