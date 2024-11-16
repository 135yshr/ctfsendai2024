package dto

import (
	"time"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
)

// ReservationResponse 予約情報レスポンス
// @Description 予約情報の詳細.
type ReservationResponse struct {
	// 予約ID
	ID string `example:"rsv_123456" json:"id"`

	// ユーザーID
	UserID string `example:"user123" json:"user_id"`

	// 予約開始時間
	StartTime time.Time `example:"2024-03-20T10:00:00Z" json:"start_time"`

	// 予約終了時間
	EndTime time.Time `example:"2024-03-20T11:00:00Z" json:"end_time"`

	// 予約ステータス
	Status string `example:"confirmed" json:"status"`

	// ユーザー情報
	User *UserResponse `json:"user"`

	// プラン情報
	Plan *PlanResponse `json:"plan"`
}

func ToReservationResponse(reservation *models.Reservation) *ReservationResponse {
	return &ReservationResponse{
		ID:        reservation.ID,
		UserID:    reservation.UserID,
		StartTime: reservation.StartTime,
		EndTime:   reservation.EndTime,
		Status:    string(reservation.Status),
		User:      ToUserResponse(reservation.User),
		Plan:      ToPlanResponse(reservation.Plan),
	}
}

type CreateReservationRequest struct {
	UserID    string    `json:"user_id"`
	PlanID    string    `json:"plan_id"`
	StartDate time.Time `json:"start_date"`
}
