package main

import (
	"github.com/gin-gonic/gin"
	"microservice/routes"
)

func main() {
	r := gin.Default()
	routes.UserRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		panic("Failed to start server")
	}
}
