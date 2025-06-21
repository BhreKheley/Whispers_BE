package routes

import (
	"github.com/BhreKheley/whispers_be/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Seat Endpoints
	r.GET("/seats", controllers.GetAllSeats)               // ✅ List semua kursi
	r.GET("/seats/booked", controllers.GetBookedSeatIDs)   // ✅ Kursi yang sudah dibooking

	// Booking Endpoints
	r.POST("/booking", controllers.CreateBooking)          // ✅ Booking baru + unggah bukti + data pemesan & tiket

	// Ticket Endpoints
	r.GET("/ticket/validate/:code", controllers.ValidateTicket) // ✅ Validasi tiket by kode (QR scan)
	r.PATCH("/ticket/checkin/:code", controllers.CheckInTicket) // ✅ Check-in manual (ubah status is_checked_in)
}
