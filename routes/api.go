package routes

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/controllers"
)

func MainRoutes(r *gin.Engine) {
	r.GET("/", controllers.Welcome)

	fp := r.Group("/fast-pay")
	{
		fp.POST("/get-link", controllers.FastPayGetLink)
	}
}
