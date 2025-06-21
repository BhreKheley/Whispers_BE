package models

import (
	"github.com/google/uuid"
)

type Seat struct {
	ID         uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	SeatCode   string        `gorm:"type:varchar(10);not null;unique"`
	Section    string        `gorm:"type:varchar(20);not null"`
	IsActive   bool          `gorm:"default:true"`
	CategoryID uuid.UUID     `gorm:"type:uuid;not null"`
	Category   SeatCategory  `gorm:"foreignKey:CategoryID"`
	Tickets    []Ticket      `gorm:"foreignKey:SeatID"`
}
