package models

import (
	"time"

	"github.com/google/uuid"
)

type CheckinLog struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	TicketID   uuid.UUID
	TiketKode  string
	ScannedAt  time.Time
	Device     string // opsional: "Front Desk A", "Mobile Scanner", dll
}
