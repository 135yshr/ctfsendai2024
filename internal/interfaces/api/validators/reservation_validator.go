package validators

import (
	"time"
)

// GetReservationsRequest 予約一覧取得リクエスト
// @Description 予約一覧を取得するためのリクエストパラメータ.
type GetReservationsRequest struct {
	// ユーザーID
	// @Example "user123"
	UserID string `binding:"required,min=3,max=50" form:"user_id"`
}

type CreateReservationRequest struct {
	UserID    string    `binding:"required" json:"user_id"`
	PlanID    string    `binding:"required" json:"plan_id"`
	StartDate time.Time `binding:"required" json:"start_date"`
}

// カスタムバリデーションメッセージ.
func GetValidationMessages() map[string]string {
	return map[string]string{
		"required": "{0}は必須項目です",
		"min":      "{0}は最低{1}文字必要です",
		"max":      "{0}は最大{1}文字までです",
		"datetime": "{0}は正しい日時形式で入力してください",
		"gtfield":  "終了時刻は開始時刻より後である必要があります",
	}
}
