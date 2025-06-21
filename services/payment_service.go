package services

import (
	"mime/multipart"
	"path/filepath"
	"strings"
)

// Validasi ekstensi file yang diperbolehkan
func IsValidPaymentProof(fileHeader *multipart.FileHeader) bool {
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	validExt := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".pdf": true,
	}
	return validExt[ext]
}
