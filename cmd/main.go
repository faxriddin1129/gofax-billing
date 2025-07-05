package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"microservice/internal/config"
	"microservice/internal/migrations"
	"microservice/pkg/bootstrap"
	"microservice/pkg/env"
	"microservice/pkg/utils"
	"time"
)

func init() {
	location, _ := time.LoadLocation("Asia/Tashkent")
	time.Local = location
}

func main() {
	// LOAD ENVIRONMENTS
	env.LoadEnv()

	// GIN DEFAULT
	r := gin.Default()

	// GIN MODE
	gin.SetMode(env.GetEnv("MODE"))

	//LOGGER
	r.Use(bootstrap.SetupErrorLogger())

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Database connection
	config.DBConnect()
	db := config.GetDB()
	utils.SetDB(db)

	// Database migrations
	if err := migrations.MigrateAll(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Routes
	config.RegisterRoutes(r)

	// Start server
	port := env.GetEnv("PROJECT_PORT")
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
