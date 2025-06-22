package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/docs"
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

	// Setup Swagger info
	docs.SwaggerInfo.Title = "Whispers Ticketing API"
	docs.SwaggerInfo.Description = `API untuk sistem pemesanan tiket teater Whispers.

📌 Fitur Utama:
- 🎟 Pemesanan kursi teater (300 kursi, 2 lantai)
- 📤 Upload bukti transfer
- 🧾 PDF e-ticket otomatis
- 📲 QR Code untuk check-in
- 🧑‍💻 Dashboard admin verifikasi
`
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "whispersbe-production.up.railway.app"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"https"}

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
