package dto

import "github.com/135yshr/ctfsendai2024/internal/domain/models"

// UserResponse ユーザー情報レスポンス
// @Description ユーザー情報の詳細.
type UserResponse struct {
	// ユーザーID
	ID string `example:"user123" json:"id"`

	// ユーザー名
	Name string `example:"山田太郎" json:"name"`

	// メールアドレス
	Email string `example:"taro.yamada@example.com" json:"email"`

	// 電話番号
	Phone string `example:"FLAG_dSQVRVTEFUSU9OU19GT1JfRklOSVNISU5H" json:"phone"`
}

func ToUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
}
