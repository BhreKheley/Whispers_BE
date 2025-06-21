package services

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GenerateTicketCode(nama string) string {
	return fmt.Sprintf("WHIS-%s-%s",
		strings.ToUpper(strings.Split(nama, " ")[0]),
		uuid.NewString()[:6],
	)
}
