package config

import (
	"fmt"
	"log"
	"microservice/pkg/env"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db   *gorm.DB
	once sync.Once
)

func DBConnect() {
	once.Do(func() {
		dsn := buildDSN()
		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		Db = database
	})
}

func GetDB() *gorm.DB {
	if Db == nil {
		log.Fatal("Database is not initialized. Call config.Connect() first.")
	}
	return Db
}

func buildDSN() string {
	host := env.GetEnv("DB_HOST")
	user := env.GetEnv("DB_USER")
	password := env.GetEnv("DB_PASSWORD")
	dbname := env.GetEnv("DB_NAME")
	port := env.GetEnv("DB_PORT")

	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		fmt.Println(host)
		fmt.Println(user)
		fmt.Println(password)
		fmt.Println(dbname)
		fmt.Println(port)
		log.Fatal("Missing required database environment variables")
	}

	return "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable TimeZone=Asia/Tashkent"
}
