package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/google/uuid"
)

func compactUUID() uuid.UUID {
	clean := strings.ReplaceAll(uuid.New().String(), "-", "")
	id, _ := uuid.Parse(clean)
	return id
}

func SeedDatabase() {
	db := config.DB

	// Cek apakah seat category sudah ada
	var count int64
	db.Model(&models.SeatCategory{}).Count(&count)
	if count > 0 {
		log.Println("🟡 Database already seeded, skipping seeder")
		return
	}

	log.Println("🚀 Seeding Seat Categories and Seats...")

	// Buat kategori
	firstFloor := models.SeatCategory{
		ID:    compactUUID(),
		Name:  "1st Floor",
		Harga: 50000,
	}
	secondFloor := models.SeatCategory{
		ID:    compactUUID(),
		Name:  "2nd Floor",
		Harga: 35000,
	}
	db.Create(&firstFloor)
	db.Create(&secondFloor)

	// Buat kursi
	var seats []models.Seat

	// A–G = 1st Floor (30 seat per row)
	for row := 'A'; row <= 'G'; row++ {
		for num := 1; num <= 30; num++ {
			code := fmt.Sprintf("%c%d", row, num)
			seats = append(seats, models.Seat{
				ID:         compactUUID(),
				SeatCode:   code,
				Section:    "1st floor",
				IsActive:   true,
				CategoryID: firstFloor.ID,
			})
		}
	}

	// H–J = 2nd Floor
	for row := 'H'; row <= 'J'; row++ {
		for num := 1; num <= 30; num++ {
			code := fmt.Sprintf("%c%d", row, num)
			seats = append(seats, models.Seat{
				ID:         compactUUID(),
				SeatCode:   code,
				Section:    "2nd floor",
				IsActive:   true,
				CategoryID: secondFloor.ID,
			})
		}
	}

	db.Create(&seats)
	log.Println("✅ Seeding complete: 2 categories, 300 seats")
}
