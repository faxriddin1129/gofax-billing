package utils

import (
	"github.com/gin-gonic/gin"
)

func RespondJson(c *gin.Context, data interface{}, code int, message string) {
	if data == nil {
		data = []string{}
	}
	c.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
}
