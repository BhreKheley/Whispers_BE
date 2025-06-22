package dto

type CheckInRequest struct {
	TiketKode string `json:"tiket_kode" binding:"required"`
}
