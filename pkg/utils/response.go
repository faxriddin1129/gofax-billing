package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, statusCode int, message map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(message)
	w.WriteHeader(statusCode)
	_, err := w.Write(res)
	if err != nil {
		panic(err)
	}
}

func RespondWithSuccess(c *gin.Context, data interface{}, code int, message string) {
	if data == nil {
		data = []string{}
	}
	c.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
}
