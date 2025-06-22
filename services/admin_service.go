package services

import (
	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"golang.org/x/crypto/bcrypt"
)

func VerifyAdminLogin(email, password string) (bool, models.Admin, error) {
	var admin models.Admin
	err := config.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return false, models.Admin{}, nil // tidak error fatal, hanya tidak ditemukan
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return false, models.Admin{}, nil
	}

	return true, admin, nil
}
