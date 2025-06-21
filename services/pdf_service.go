package services

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func GenerateETicketPDF(nama string, kode string, qrPath string, outputPath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 10, "Whispers Theater E-Ticket")
	pdf.Ln(20)

	pdf.SetFont("Arial", "", 14)
	pdf.Cell(0, 10, fmt.Sprintf("Nama: %s", nama))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Kode Tiket: %s", kode))
	pdf.Ln(10)

	if qrPath != "" {
		pdf.Image(qrPath, 10, 60, 60, 60, false, "", 0, "")
	}

	return pdf.OutputFileAndClose(outputPath)
}
