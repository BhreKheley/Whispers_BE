package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/BhreKheley/whispers_be/config"
	"github.com/BhreKheley/whispers_be/models"
	"github.com/jordan-wright/email"
)

func SendTicketsEmail(booking models.Booking) error {
	var tickets []models.Ticket
	if err := config.DB.Where("booking_id = ?", booking.ID).Find(&tickets).Error; err != nil {
		return fmt.Errorf("failed to fetch tickets: %v", err)
	}

	e := email.NewEmail()
	e.From = fmt.Sprintf("Whispers Tiket <%s>", os.Getenv("EMAIL_SENDER"))
	e.To = []string{booking.EmailPemesan}
	e.Subject = "E-Ticket Whispers 🎭"
	e.Text = []byte(fmt.Sprintf("Halo %s,\n\nBerikut adalah e-ticket Anda. Mohon tunjukkan QR Code saat check-in di venue.\n\nSalam,\nTim Whispers", booking.NamaPemesan))

	// Lampirkan PDF tiket
	for _, t := range tickets {
		path := filepath.Join("tickets", filepath.Base(t.PDFPath))
		if _, err := e.AttachFile(path); err != nil {
			log.Printf("❌ Gagal attach PDF: %v\n", err)
		}
	}

	// Kirim email
	smtpAddr := "smtp.gmail.com:587"
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_SENDER"), os.Getenv("EMAIL_PASSWORD"), "smtp.gmail.com")
	err := e.Send(smtpAddr, auth)
	if err != nil {
		return fmt.Errorf("gagal mengirim email: %v", err)
	}

	log.Printf("📧 Email e-ticket terkirim ke %s\n", booking.EmailPemesan)
	return nil
}
