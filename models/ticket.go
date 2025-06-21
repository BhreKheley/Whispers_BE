package models

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	BookingID      uuid.UUID `gorm:"type:uuid;not null"`
	SeatID         uuid.UUID `gorm:"type:uuid;not null"`
	NamaPemegang   string    `gorm:"type:varchar(100);not null"`
	EmailPemegang  string    `gorm:"type:varchar(100);not null"`
	HPPemegang     string    `gorm:"type:varchar(20);not null"`
	TiketKode      string    `gorm:"type:varchar(100);not null;unique"`
	QRPath         string    `gorm:"type:text"` // ✅ Tambahkan
	PDFPath        string    `gorm:"type:text"` // ✅ Tambahkan
	IsCheckedIn    bool      `gorm:"default:false"`

	Booking Booking `gorm:"foreignKey:BookingID"`
	Seat    Seat    `gorm:"foreignKey:SeatID"`
}
