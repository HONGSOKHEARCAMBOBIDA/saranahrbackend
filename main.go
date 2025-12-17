package main

import (
	"HRbackend/config"
	"HRbackend/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Create Gin router
	r := gin.Default()

	// Apply CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://192.168.39.193:5173"}, // your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Set up routes
	routes.SetupRoutes(r)

	// Start server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
