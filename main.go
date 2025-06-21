package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/routes"
	"github.com/BhreKheley/whispers_be/utils"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load (ignored in production)")
	}

	// Init database
	config.InitDB()

	// Panggil Seeder
	utils.SeedDatabase()

	// Setup Gin
	r := gin.Default()

	// Optional: Allow CORS for React FE
	r.Use(cors.Default())

	// Serve uploads folder (for bukti, QR, PDF, etc.)
	r.Static("/uploads", "./uploads")
	r.Static("/qrcodes", "./qrcodes")
	r.Static("/tickets", "./tickets")

	// Register routes
	routes.SetupRoutes(r)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Running server on port", port)
	r.Run(":" + port)
}
