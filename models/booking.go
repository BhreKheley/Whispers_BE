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
	MetodePembayaran string    `gorm:"type:varchar(20);not null"`         // qris / transfer
	BuktiTransfer    string    `gorm:"type:text"`                         // path file upload
	Status           string    `gorm:"type:varchar(20);default:'pending'"` // pending, waiting-verification, paid, rejected
	TotalHarga       int       `gorm:"not null"`                          // auto dihitung dari seat
	CreatedAt        time.Time `gorm:"autoCreateTime"`

	Tickets []Ticket `gorm:"foreignKey:BookingID"`
}
