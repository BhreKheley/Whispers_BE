package services

import (
	"time"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/google/uuid"
)

func CheckInTicket(kode string) (bool, error) {
	var ticket models.Ticket

	err := config.DB.Where("tiket_kode = ?", kode).First(&ticket).Error
	if err != nil {
		return false, err
	}

	if ticket.IsCheckedIn {
		return false, nil
	}

	ticket.IsCheckedIn = true
	if err := config.DB.Save(&ticket).Error; err != nil {
		return false, err
	}

	// ⬇️ Tambahkan log check-in
	log := models.CheckinLog{
		ID:         uuid.New(),
		TicketID:   ticket.ID,
		TiketKode:  ticket.TiketKode,
		ScannedAt:  time.Now(),
		Device:     "scanner_default", // bisa disesuaikan nanti
	}
	config.DB.Create(&log)

	return true, nil
}

