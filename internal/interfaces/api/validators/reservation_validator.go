package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// GetReservationsRequest 予約一覧取得リクエスト
// @Description 予約一覧を取得するためのリクエストパラメータ.
type GetReservationsRequest struct {
	// ユーザーID
	UserID string `binding:"required,min=3,max=50" form:"user_id"`
}

// CreateReservationRequest 予約作成リクエスト
// @Description 新規予約を作成するためのリクエストパラメータ.
type CreateReservationRequest struct {
	// ユーザーID
	UserID string `binding:"required" json:"user_id"`

	// プランID
	PlanID string `binding:"required" json:"plan_id"`

	// 予約開始日時
	StartDate time.Time `binding:"required,future" json:"start_date"`
}

// カスタムバリデーションメッセージ.
func GetValidationMessages() map[string]string {
	return map[string]string{
		"required": "{0}は必須項目です",
		"min":      "{0}は最低{1}文字必要です",
		"max":      "{0}は最大{1}文字までです",
		"datetime": "{0}は正しい日時形式で入力してください",
		"gtfield":  "終了時刻は開始時刻より後である必要があります",
		"future":   "{0}は現在時刻より後の時間を指定してください",
	}
}

// カスタムバリデーションを登録する関数を追加.
func RegisterCustomValidations(v *validator.Validate) {
	_ = v.RegisterValidation("future", validateFutureDate)
}

// 未来の日付かどうかを検証するバリデーション関数.
func validateFutureDate(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	return date.After(time.Now())
}
