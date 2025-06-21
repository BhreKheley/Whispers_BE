package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/BhreKheley/whispers_be/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TicketInput struct {
	SeatID         uuid.UUID `json:"seat_id"`
	NamaPemegang   string    `json:"nama_pemegang"`
	EmailPemegang  string    `json:"email_pemegang"`
	HpPemegang     string    `json:"hp_pemegang"`
}

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	var tiketInputs []TicketInput

	// Ambil data pemesan
	booking.ID = uuid.New()
	booking.NamaPemesan = c.PostForm("nama_pemesan")
	booking.EmailPemesan = c.PostForm("email_pemesan")
	booking.NoHP = c.PostForm("no_hp")
	booking.MetodePembayaran = c.PostForm("metode_pembayaran")
	booking.Status = "pending"
	booking.CreatedAt = time.Now()

	// Upload bukti transfer
	file, err := c.FormFile("bukti_transfer")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bukti transfer wajib diunggah"})
		return
	}
	filename := fmt.Sprintf("%s_%s", booking.ID.String(), file.Filename)
	filePath := filepath.Join("uploads", filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file bukti"})
		return
	}
	booking.BuktiTransfer = filename

	// Ambil dan parse tiket JSON
	tiketJson := c.PostForm("tickets")
	if err := json.Unmarshal([]byte(tiketJson), &tiketInputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format tiket tidak valid"})
		return
	}

	// Validasi seat dan hitung total harga
	var totalHarga int
	var tiketList []models.Ticket
	var seatIDs []uuid.UUID

	for _, t := range tiketInputs {
		// Validasi seat belum dibooking
		available, err := services.IsSeatAvailable(t.SeatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengecek seat"})
			return
		}
		if !available {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Seat %s sudah dibooking", t.SeatID)})
			return
		}
		seatIDs = append(seatIDs, t.SeatID)
	}

	totalHarga, err = services.CalculateTotalHarga(seatIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung total harga"})
		return
	}
	booking.TotalHarga = totalHarga

	// Simpan booking
	if err := config.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan booking"})
		return
	}

	// Proses dan simpan tiket
	for _, t := range tiketInputs {
		ticketCode := services.GenerateTicketCode(t.NamaPemegang)

		qrPath, err := services.GenerateQRCode(ticketCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal generate QR code"})
			return
		}

		pdfPath := fmt.Sprintf("tickets/%s.pdf", ticketCode)
		err = services.GenerateETicketPDF(t.NamaPemegang, ticketCode, qrPath, pdfPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal generate PDF e-ticket"})
			return
		}

		tiket := models.Ticket{
			ID:             uuid.New(),
			BookingID:      booking.ID,
			SeatID:         t.SeatID,
			NamaPemegang:   t.NamaPemegang,
			EmailPemegang:  t.EmailPemegang,
			HPPemegang:     t.HpPemegang,
			TiketKode:      ticketCode,
			QRPath:         qrPath,
			PDFPath:        pdfPath,
			IsCheckedIn:    false,
		}
		tiketList = append(tiketList, tiket)

		if err := config.DB.Create(&tiket).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan tiket"})
			return
		}
	}

	// Sukses
	c.JSON(http.StatusOK, gin.H{
		"message":       "Booking berhasil disimpan",
		"booking_id":    booking.ID,
		"jumlah_tiket":  len(tiketList),
		"total_harga":   booking.TotalHarga,
	})
}
