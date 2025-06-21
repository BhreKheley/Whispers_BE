package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	NamaPemesan      string    `gorm:"type:varchar(100);not null"`
	EmailPemesan     string    `gorm:"type:varchar(100);not null"`
	NoHP             string    `gorm:"type:varchar(20);not null"`
	MetodePembayaran string    `gorm:"type:varchar(20);not null"`
	BuktiTransfer    string    `gorm:"type:text"`
	Status           string    `gorm:"type:varchar(20);default:'pending'"`
	TotalHarga       int       `gorm:"not null"` // ✅ Tambahkan ini
	CreatedAt        time.Time `gorm:"autoCreateTime"`

	Tickets []Ticket `gorm:"foreignKey:BookingID"`
}
