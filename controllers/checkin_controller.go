package controllers

import (
	"net/http"

	"github.com/BhreKheley/whispers_be/dto"
	"github.com/BhreKheley/whispers_be/services"
	"github.com/gin-gonic/gin"
)

// CheckInTicket godoc
// @Summary Check-in tiket menggunakan kode QR
// @Tags Ticket
// @Accept json
// @Produce json
// @Param payload body dto.CheckInRequest true "Kode Tiket"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} map[string]interface{}
// @Router /ticket/checkin [post]
func CheckInTicket(c *gin.Context) {
	var req dto.CheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := services.CheckInTicket(req.TiketKode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tiket tidak ditemukan"})
		return
	}

	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tiket sudah digunakan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Check-in berhasil. Selamat menonton!"})
}
