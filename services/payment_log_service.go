package services

import (
	"time"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/google/uuid"
)

func LogPaymentAction(bookingID uuid.UUID, action, note, filePath string) error {
	logEntry := models.PaymentLog{
		ID:        uuid.New(),
		BookingID: bookingID,
		Action:    action,     // "uploaded", "approved", "rejected"
		Note:      note,
		FilePath:  filePath,
		CreatedAt: time.Now(),
	}
	return config.DB.Create(&logEntry).Error
}
