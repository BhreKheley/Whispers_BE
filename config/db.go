package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BhreKheley/whispers_be/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load(".env")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Println("No .env file found or failed to load (ignored in production)")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database:%v", err)
	}
	fmt.Println("Database connection established")

	// ✅ Auto migrate semua tabel
	err = DB.AutoMigrate(
		&models.SeatCategory{},
		&models.Seat{},
		&models.Booking{},
		&models.Ticket{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}
	fmt.Println("AutoMigrate completed ✅")
}
