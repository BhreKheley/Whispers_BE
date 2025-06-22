package models

import (
	"github.com/google/uuid"
)

type SeatCategory struct {
	ID    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name  string    `gorm:"type:varchar(50);not null;unique"` // Contoh: VIP, 1st Floor
	Harga int       `gorm:"not null"`                         // Harga dalam rupiah

	Seats []Seat `gorm:"foreignKey:CategoryID"`
}
