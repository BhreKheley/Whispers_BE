package services

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"
)

// Generate QR untuk tiket berdasarkan kode unik
func GenerateQRCode(tiketKode string) (string, error) {
	dir := "qrcodes"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	filePath := fmt.Sprintf("%s/%s.png", dir, tiketKode)
	err := qrcode.WriteFile(tiketKode, qrcode.Medium, 256, filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
