package controllers

import (
	"net/http"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/dto"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/BhreKheley/whispers_be/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AdminLogin godoc
// @Summary Login admin (sementara tanpa JWT)
// @Description Autentikasi admin menggunakan email dan password
// @Tags Admin
// @Accept json
// @Produce json
// @Param credentials body dto.AdminLoginRequest true "Email dan Password Admin"
// @Success 200 {object} map[string]interface{} "Login berhasil"
// @Failure 400 {object} map[string]interface{} "Request tidak valid"
// @Failure 401 {object} map[string]interface{} "Kredensial salah"
// @Failure 500 {object} map[string]interface{} "Server error"
// @Router /admin/login [post]
func AdminLogin(c *gin.Context) {
	var req dto.AdminLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, admin, err := services.VerifyAdminLogin(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"admin": gin.H{
			"id":    admin.ID,
			"name":  admin.Name,
			"email": admin.Email,
		},
	})
}

// ApproveBooking godoc
// @Summary Verifikasi dan approve booking
// @Description Mengubah status booking menjadi 'paid' dan menyimpan log pembayaran
// @Tags Admin
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} map[string]interface{} "Booking berhasil di-approve"
// @Failure 400 {object} map[string]interface{} "ID booking tidak valid"
// @Failure 404 {object} map[string]interface{} "Booking tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menyimpan perubahan"
// @Router /admin/approve/{id} [patch]
func ApproveBooking(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	var booking models.Booking
	if err := config.DB.First(&booking, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking tidak ditemukan"})
		return
	}

	booking.Status = "paid"
	if err := config.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate status"})
		return
	}

	_ = services.LogPaymentAction(booking.ID, "approved", "Approved by admin", booking.BuktiTransfer)

	c.JSON(http.StatusOK, gin.H{"message": "Booking approved"})
}

// RejectBooking godoc
// @Summary Tolak dan reject booking
// @Description Mengubah status booking menjadi 'rejected' dan menyimpan log pembayaran
// @Tags Admin
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} map[string]interface{} "Booking berhasil ditolak"
// @Failure 400 {object} map[string]interface{} "ID booking tidak valid"
// @Failure 404 {object} map[string]interface{} "Booking tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menyimpan perubahan"
// @Router /admin/reject/{id} [delete]
func RejectBooking(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	var booking models.Booking
	if err := config.DB.First(&booking, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking tidak ditemukan"})
		return
	}

	booking.Status = "rejected"
	if err := config.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate status"})
		return
	}

	_ = services.LogPaymentAction(booking.ID, "rejected", "Ditolak oleh admin", booking.BuktiTransfer)

	c.JSON(http.StatusOK, gin.H{"message": "Booking ditolak"})
}
