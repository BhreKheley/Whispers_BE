package models

import (
	"github.com/google/uuid"
)

type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Email    string    `gorm:"type:varchar(100);not null;unique"`
	Password string    `gorm:"type:varchar(100);not null"` // hash (gunakan bcrypt)
	Name     string    `gorm:"type:varchar(50)"`
}
