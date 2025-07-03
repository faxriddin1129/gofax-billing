package routes

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/middleware"
)

func MainRoutes(r *gin.Engine) {
	r.GET("/ping", middleware.AuthMiddleware(), func(c *gin.Context) {
		userID := c.GetUint("userID")
		c.JSON(200, gin.H{
			"message": "Ping server",
			"user_id": userID,
		})
	})
}
