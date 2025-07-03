package routes

import (
	"github.com/gin-gonic/gin"
	"microservice/controllers"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
		users.POST("/", controllers.CreateUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
}
