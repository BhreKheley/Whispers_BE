package services

import (
	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/google/uuid"
)

func IsSeatAvailable(seatID uuid.UUID) (bool, error) {
	var count int64
	err := config.DB.Model(&models.Ticket{}).Where("seat_id = ?", seatID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func CalculateTotalHarga(seatIDs []uuid.UUID) (int, error) {
	var seats []models.Seat
	err := config.DB.Preload("Category").Where("id IN ?", seatIDs).Find(&seats).Error
	if err != nil {
		return 0, err
	}

	total := 0
	for _, seat := range seats {
		total += seat.Category.Harga
	}
	return total, nil
}

func SaveBookingWithTickets(
	booking *models.Booking,
	tickets []models.Ticket,
) error {
	tx := config.DB.Begin()

	if err := tx.Create(booking).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, ticket := range tickets {
		if err := tx.Create(&ticket).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
