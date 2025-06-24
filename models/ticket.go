package models

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	BookingID   uuid.UUID `gorm:"type:uuid;not null"`
	SeatID      uuid.UUID `gorm:"type:uuid;not null;unique"`
	TiketKode   string    `gorm:"type:varchar(100);not null;unique"`
	QRPath      string    `gorm:"type:text"`  // path file QR code
	PDFPath     string    `gorm:"type:text"`  // path file PDF e-ticket
	IsCheckedIn bool      `gorm:"default:false"`

	Booking Booking `gorm:"foreignKey:BookingID"`
	Seat    Seat    `gorm:"foreignKey:SeatID"`
}
