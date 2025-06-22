package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/routes"
	"github.com/BhreKheley/whispers_be/utils"

	_ "github.com/BhreKheley/whispers_be/docs" // docs init (penting!)
)


// @title Whispers Ticketing API
// @version 1.0
// @description Sistem pemesanan tiket teater Whispers.
// @contact.name Kheleyome
// @contact.email kheleyome1@gmail.com
// @host localhost:8080
// @BasePath /
func main() {
	// Load environment variables
	config.LoadEnv()

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

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Running server on port", port)
	r.Run(":" + port)
}
