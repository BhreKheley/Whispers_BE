package models

import (
	"github.com/google/uuid"
)

type Seat struct {
	ID         uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	SeatCode   string       `gorm:"type:varchar(10);not null;unique"` // Contoh: A1, B10
	Section    string       `gorm:"type:varchar(20);not null"`        // 1st floor, 2nd floor
	IsActive   bool         `gorm:"default:true"`                     // False = dibooking
	CategoryID uuid.UUID    `gorm:"type:uuid;not null"`
	Category   SeatCategory `gorm:"foreignKey:CategoryID"`

	Tickets []Ticket `gorm:"foreignKey:SeatID"`
}
