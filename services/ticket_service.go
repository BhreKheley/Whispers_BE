package services

import (
	"fmt"
	"github.com/google/uuid"
)

// Versi final, tanpa input nama
func GenerateTicketCode() string {
	return fmt.Sprintf("WHIS-%s", uuid.NewString()[:8])
}
