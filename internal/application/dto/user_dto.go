package dto

import "github.com/135yshr/ctfsendai2024/internal/domain/models"

// UserResponse ユーザー情報レスポンス
// @Description ユーザー情報の詳細.
type UserResponse struct {
	// ユーザーID
	// @Example "user123"
	ID string `example:"user123" json:"id"`

	// ユーザー名
	// @Example "山田太郎"
	Name string `example:"山田太郎" json:"name"`

	// メールアドレス
	// @Example "taro.yamada@example.com"
	Email string `example:"taro.yamada@example.com" json:"email"`

	// 電話番号
	// @Example "090-1234-5678"
	Phone string `example:"090-1234-5678" json:"phone"`
}

func ToUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
}
