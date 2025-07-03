package payme

import (
	"github.com/gin-gonic/gin"
)

func PaymeRoutes(r *gin.Engine) {
	routes := r.Group("/payme")
	{
		routes.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Payme working",
			})
		})
	}
}
