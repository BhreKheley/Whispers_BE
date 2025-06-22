package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentLog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	BookingID uuid.UUID
	Action    string    // uploaded, approved, rejected
	Note      string    // opsional: alasan ditolak, catatan admin
	FilePath  string    // path file bukti (jika ada)
	CreatedAt time.Time
}
