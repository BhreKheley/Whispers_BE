package controllers

import (
	"net/http"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/gin-gonic/gin"
)

func ValidateTicket(c *gin.Context) {
	code := c.Param("code")
	var ticket models.Ticket

	err := config.DB.Preload("Seat.Category").
		Where("tiket_kode = ?", code).
		First(&ticket).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tiket tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Tiket valid",
		"ticket_data": ticket,
	})
}

func CheckInTicket(c *gin.Context) {
	code := c.Param("code")
	var ticket models.Ticket

	if err := config.DB.Where("tiket_kode = ?", code).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tiket tidak ditemukan"})
		return
	}

	if ticket.IsCheckedIn {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tiket sudah check-in sebelumnya"})
		return
	}

	ticket.IsCheckedIn = true
	if err := config.DB.Save(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update status check-in"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tiket berhasil check-in"})
}
