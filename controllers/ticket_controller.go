package controllers

import (
	"net/http"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/gin-gonic/gin"
)

// ValidateTicket godoc
// @Summary Validasi tiket berdasarkan kode
// @Tags Ticket
// @Produce json
// @Param code path string true "Kode Tiket"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /ticket/validate/{code} [get]
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
