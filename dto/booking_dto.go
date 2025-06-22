package dto

// Digunakan saat user mengisi data pemesanan tiket
type CreateBookingRequest struct {
	NamaPemesan      string   `json:"nama_pemesan" binding:"required"`
	EmailPemesan     string   `json:"email_pemesan" binding:"required,email"`
	NoHP             string   `json:"no_hp" binding:"required"`
	MetodePembayaran string   `json:"metode_pembayaran" binding:"required,oneof=qris transfer"`
	SeatIDs          []string `json:"seat_ids" binding:"required,dive,uuid"`
}
