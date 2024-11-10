package dto

import (
	"time"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
)

// ReservationResponse 予約情報レスポンス
// @Description 予約情報の詳細
type ReservationResponse struct {
	// 予約ID
	// @Example "rsv_123456"
	ID string `json:"id" example:"rsv_123456"`

	// ユーザーID
	// @Example "user123"
	UserID string `json:"user_id" example:"user123"`

	// 予約開始時間
	// @Example "2024-03-20T10:00:00Z"
	StartTime time.Time `json:"start_time" example:"2024-03-20T10:00:00Z"`

	// 予約終了時間
	// @Example "2024-03-20T11:00:00Z"
	EndTime time.Time `json:"end_time" example:"2024-03-20T11:00:00Z"`

	// 予約ステータス
	// @Example "confirmed"
	Status string `json:"status" example:"confirmed"`

	// ユーザー情報
	// @Example {"id": "user123", "name": "山田太郎"}
	User *UserResponse `json:"user"`

	// プラン情報
	// @Example {"id": "plan123", "name": "スタンダードプラン"}
	Plan *PlanResponse `json:"plan"`
}

func ToReservationResponse(reservation *models.Reservation) *ReservationResponse {
	return &ReservationResponse{
		ID:        reservation.ID,
		UserID:    reservation.UserID,
		StartTime: reservation.StartTime,
		EndTime:   reservation.EndTime,
		Status:    reservation.Status,
		User:      ToUserResponse(&reservation.User),
		Plan:      ToPlanResponse(&reservation.Plan),
	}
}
