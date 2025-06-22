package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/dto"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/BhreKheley/whispers_be/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateBooking godoc
// @Summary Buat booking baru
// @Tags Booking
// @Accept json
// @Produce json
// @Param booking body dto.CreateBookingRequest true "Data Booking"
// @Success 201 {object} map[string]interface{}
// @Failure 400,500 {object} map[string]interface{}
// @Router /booking [post]
func CreateBooking(c *gin.Context) {
	var req dto.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Format JSON tidak valid",
			"details": err.Error(),
		})
		return
	}

	if req.NamaPemesan == "" || req.EmailPemesan == "" || req.NoHP == "" || req.MetodePembayaran == "" || len(req.SeatIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Semua field wajib diisi: nama_pemesan, email_pemesan, no_hp, metode_pembayaran, dan seat_ids",
		})
		return
	}

	var seatIDs []uuid.UUID
	for _, s := range req.SeatIDs {
		id, err := uuid.Parse(s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ID seat tidak valid (bukan UUID)",
				"details": fmt.Sprintf("Seat ID: %s, Error: %s", s, err.Error()),
			})
			return
		}
		seatIDs = append(seatIDs, id)
	}

	tx := config.DB.Begin()

	var availableSeats []models.Seat
	if err := tx.Preload("Category").
		Where("id IN ? AND is_active = true", seatIDs).
		Find(&availableSeats).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal cek ketersediaan kursi"})
		return
	}

	if len(availableSeats) != len(seatIDs) {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Beberapa kursi sudah dibooking oleh orang lain",
		})
		return
	}

	// Hitung total harga
	total := 0
	for _, seat := range availableSeats {
		total += seat.Category.Harga
	}

	// Simpan Booking
	booking := models.Booking{
		ID:               uuid.New(),
		NamaPemesan:      req.NamaPemesan,
		EmailPemesan:     req.EmailPemesan,
		NoHP:             req.NoHP,
		MetodePembayaran: req.MetodePembayaran,
		Status:           "pending",
		CreatedAt:        time.Now(),
		TotalHarga:       total,
	}
	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan booking"})
		return
	}

	for _, seat := range availableSeats {
		tiketKode := services.GenerateTicketCode()
		qrPath, err := services.GenerateQRCode(tiketKode)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal generate QR code"})
			return
		}

		pdfPath := fmt.Sprintf("tickets/%s.pdf", tiketKode)
		if err := services.GenerateETicketPDF(booking.NamaPemesan, tiketKode, qrPath, pdfPath); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal generate PDF tiket"})
			return
		}

		ticket := models.Ticket{
			ID:          uuid.New(),
			BookingID:   booking.ID,
			SeatID:      seat.ID,
			TiketKode:   tiketKode,
			QRPath:      qrPath,
			PDFPath:     pdfPath,
			IsCheckedIn: false,
		}
		if err := tx.Create(&ticket).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan tiket"})
			return
		}

		// Update seat jadi tidak aktif
		if err := tx.Model(&models.Seat{}).
			Where("id = ?", seat.ID).
			Update("is_active", false).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update status kursi"})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyelesaikan transaksi"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Booking berhasil",
		"booking_id":  booking.ID,
		"total_harga": total,
		"seat_count":  len(seatIDs),
	})
}

// UploadBuktiTransfer godoc
// @Summary Upload bukti transfer
// @Tags Booking
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Booking ID"
// @Param bukti_transfer formData file true "File Bukti Transfer"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404,500 {object} map[string]interface{}
// @Router /booking/{id}/upload [post]
func UploadBuktiTransfer(c *gin.Context) {
	bookingID := c.Param("id")
	id, err := uuid.Parse(bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Booking ID tidak valid (UUID expected)",
			"details": err.Error(),
		})
		return
	}

	file, err := c.FormFile("bukti_transfer")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bukti transfer tidak ditemukan dalam request",
			"details": err.Error(),
		})
		return
	}

	if !services.IsValidPaymentProof(file) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Format bukti transfer tidak valid. Hanya mendukung jpg, png, pdf",
		})
		return
	}

	filename := fmt.Sprintf("%s_%s", id.String(), file.Filename)
	filePath := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Gagal menyimpan file bukti transfer",
			"details": err.Error(),
		})
		return
	}

	var booking models.Booking
	if err := config.DB.First(&booking, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Booking tidak ditemukan",
			"details": err.Error(),
		})
		return
	}

	booking.BuktiTransfer = filename
	booking.Status = "waiting-verification"
	if err := config.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Gagal menyimpan perubahan pada booking",
			"details": err.Error(),
		})
		return
	}

	_ = services.LogPaymentAction(booking.ID, "uploaded", "User uploaded proof", filename)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Bukti transfer berhasil diupload",
		"booking_id": booking.ID,
		"file":       filename,
	})
}
