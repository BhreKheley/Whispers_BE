package dto

type UploadBuktiRequest struct {
	BookingID string `form:"booking_id" binding:"required,uuid"`
}
