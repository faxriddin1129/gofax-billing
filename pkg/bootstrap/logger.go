package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"time"
)

func SetupErrorLogger() gin.HandlerFunc {
	today := time.Now().Format("2006-01-02")
	logFile := filepath.Join("storage", "logs", today+"_error.log")

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Log fayl ochilmadi: %v", err))
	}

	return gin.CustomRecoveryWithWriter(io.MultiWriter(f, os.Stderr), func(c *gin.Context, recovered interface{}) {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Internal Server Error",
		})
	})
}
