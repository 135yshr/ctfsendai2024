package dto

import "github.com/135yshr/ctfsendai2024/internal/domain/models"

// UserResponse ユーザー情報レスポンス
// @Description ユーザー情報の詳細
type UserResponse struct {
	// ユーザーID
	// @Example "user123"
	ID string `json:"id" example:"user123"`

	// ユーザー名
	// @Example "山田太郎"
	Name string `json:"name" example:"山田太郎"`

	// メールアドレス
	// @Example "taro.yamada@example.com"
	Email string `json:"email" example:"taro.yamada@example.com"`

	// 電話番号
	// @Example "090-1234-5678"
	Phone string `json:"phone" example:"090-1234-5678"`
}

func ToUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
}
