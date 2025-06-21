package models

import (
	"github.com/google/uuid"
)

type SeatCategory struct {
	ID    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name  string    `gorm:"type:varchar(50);not null;unique"`
	Harga int       `gorm:"not null"`
	Seats []Seat    `gorm:"foreignKey:CategoryID"`
}
