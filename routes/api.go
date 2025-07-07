package routes

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/controllers"
)

func MainRoutes(r *gin.Engine) {

	r.GET("/api", controllers.Welcome)

	fp := r.Group("api/fast-pay")
	{
		fp.POST("/get-link", controllers.FastPayGetLink)
		fp.POST("/get-link-by-card", controllers.FastPayByCardGetLink)
	}

	r.GET("/api/octo/notify", controllers.Welcome)
}
