package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gofax-billing/internal/config"
	"gofax-billing/internal/migrations"
	"gofax-billing/pkg/bootstrap"
	"gofax-billing/pkg/env"
	"gofax-billing/pkg/utils"
	"log"
	"time"
)

func init() {
	location, _ := time.LoadLocation("Asia/Tashkent")
	time.Local = location
}

func main() {
	// LOAD ENVIRONMENTS
	env.LoadEnv()

	// GIN MODE
	gin.SetMode(env.GetEnv("GIN_MODE"))

	// GIN DEFAULT
	r := gin.Default()

	//LOGGER
	r.Use(bootstrap.RequestResponseLogger())

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
