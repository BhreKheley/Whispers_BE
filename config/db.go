package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BhreKheley/whispers_be/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in environment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("✅ Connected to database")

	// ✅ Auto migrate semua tabel
	err = DB.AutoMigrate(
		&models.SeatCategory{},
		&models.Seat{},
		&models.Booking{},
		&models.Ticket{},
		&models.Admin{},
		&models.CheckinLog{},
		&models.PaymentLog{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}
	fmt.Println("AutoMigrate completed ✅")
}
