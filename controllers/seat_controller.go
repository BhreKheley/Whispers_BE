package controllers

import (
	"net/http"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/gin-gonic/gin"
)

// GET /seats → semua kursi + kategori
func GetAllSeats(c *gin.Context) {
	var seats []models.Seat

	err := config.DB.Preload("Category").Find(&seats).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kursi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"seats": seats})
}

func GetBookedSeatIDs(c *gin.Context) {
	var tickets []models.Ticket
	err := config.DB.Select("seat_id").Find(&tickets).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil tiket"})
		return
	}

	var seatIDs []string
	for _, t := range tickets {
		seatIDs = append(seatIDs, t.SeatID.String())
	}

	c.JSON(http.StatusOK, gin.H{"booked_seat_ids": seatIDs})
}
