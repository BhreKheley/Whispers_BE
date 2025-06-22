package controllers

import (
	"net/http"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/gin-gonic/gin"
)

// GetCheckinLogs godoc
// @Summary Ambil log check-in
// @Tags Admin
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /admin/logs/checkin [get]
func GetCheckinLogs(c *gin.Context) {
	var logs []models.CheckinLog
	if err := config.DB.Order("scanned_at desc").Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data log"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"logs": logs})
}
