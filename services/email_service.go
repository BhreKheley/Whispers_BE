package services

import (
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) error {
	from := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	addr := "smtp.gmail.com:587" // contoh Gmail

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	return smtp.SendMail(addr, auth, from, []string{to}, []byte(msg))
}
