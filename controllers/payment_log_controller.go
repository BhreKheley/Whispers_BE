package controllers

import (
	"net/http"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetPaymentLogs godoc
// @Summary Ambil log pembayaran berdasarkan booking ID
// @Tags Admin
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,500 {object} map[string]interface{}
// @Router /admin/logs/payment/{id} [get]
func GetPaymentLogs(c *gin.Context) {
	bookingIDStr := c.Param("id")
	bookingID, err := uuid.Parse(bookingIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	var logs []models.PaymentLog
	if err := config.DB.Where("booking_id = ?", bookingID).Order("created_at desc").Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data log"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"logs": logs})
}
