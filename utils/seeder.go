package utils

import (
	"fmt"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/google/uuid"
)

func SeedDatabase() {
	// Cek jika sudah ada kategori
	var count int64
	config.DB.Model(&models.SeatCategory{}).Count(&count)
	if count > 0 {
		fmt.Println("🚫 Seeder sudah pernah dijalankan.")
		return
	}

	// Buat kategori
	firstFloor := models.SeatCategory{
		ID:    uuid.New(),
		Name:  "1st floor",
		Harga: 50000,
	}
	secondFloor := models.SeatCategory{
		ID:    uuid.New(),
		Name:  "2nd floor",
		Harga: 35000,
	}
	config.DB.Create(&firstFloor)
	config.DB.Create(&secondFloor)

	// Buat 300 kursi A1–J30
	rows := "ABCDEFGHIJ"
	for _, row := range rows {
		for seatNum := 1; seatNum <= 30; seatNum++ {
			section := ""
			if seatNum <= 10 {
				section = "Left"
			} else if seatNum <= 20 {
				section = "Center"
			} else {
				section = "Right"
			}

			seat := models.Seat{
				ID:         uuid.New(),
				SeatCode:   fmt.Sprintf("%c%d", row, seatNum),
				Section:    section,
				IsActive:   true,
				CategoryID: firstFloor.ID,
			}
			if row >= 'H' {
				seat.CategoryID = secondFloor.ID
			}
			config.DB.Create(&seat)
		}
	}

	fmt.Println("✅ Seeder selesai dijalankan")
}
