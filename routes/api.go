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

	octo := r.Group("/api/octo")
	{
		octo.POST("/notify", controllers.OctoShopApiNotify)
	}

	payme := r.Group("/api/payme")
	{
		payme.POST("/notify", controllers.OctoShopApiNotify)
	}

}
