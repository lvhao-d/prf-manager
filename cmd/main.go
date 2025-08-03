package main

import (
	"log"
	"os"
	"prf-manager/internal/infrastructure/gorm"
	router "prf-manager/internal/interface/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := gorm.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connection established successfully")
	}

	r := gin.Default()
	router.SetupRoutes(r, db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set in .env
	}
	r.Run(":" + port)
}
