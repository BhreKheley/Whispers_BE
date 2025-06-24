package routes

import (
	"github.com/BhreKheley/whispers_be/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// ===== 🎫 BOOKING =====
	r.POST("/booking", controllers.CreateBooking)                  // Booking baru
	r.POST("/booking/upload/:id", controllers.UploadBuktiTransfer) // Upload bukti transfer

	// ===== 👨‍💼 ADMIN LOGIN + VERIFIKASI =====
	r.POST("/admin/login", controllers.AdminLogin)                 // Login admin
	r.PATCH("/admin/approve/:id", controllers.ApproveBooking)      // Approve booking
	r.PATCH("/admin/reject/:id", controllers.RejectBooking)        // Reject booking

	// ===== 🪑 SEAT =====
	r.GET("/seats", controllers.GetAllSeats)                       // Semua kursi + kategori
	r.GET("/seats/booked", controllers.GetBookedSeatIDs)           // ID kursi yang sudah dibooking

	// ===== 📥 QR VALIDATION =====
	r.GET("/ticket/validate/:code", controllers.ValidateTicket)    // Validasi tiket dari kode QR

	// ===== ✅ CHECKIN (Frontend) =====
	r.POST("/ticket/checkin", controllers.CheckInTicket)           // QR scan manual FE

	// ===== 📜 LOGS =====
	r.GET("/logs/checkin", controllers.GetCheckinLogs)             // Log scan tiket
	r.GET("/logs/payment/:id", controllers.GetPaymentLogs)         // Log status bukti pembayaran per booking

	// ===== 🧪 DEBUG: LIST ALL ROUTES =====
	r.GET("/list-routes", func(c *gin.Context) {
		routes := r.Routes()
		var list []map[string]string
		for _, route := range routes {
			list = append(list, map[string]string{
				"method": route.Method,
				"path":   route.Path,
			})
		}
		c.JSON(200, gin.H{"routes": list})
	})
}
