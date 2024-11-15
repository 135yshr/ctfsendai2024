package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters/response"
)

// UserPresenter ユーザー情報のプレゼンテーション層のインターフェース.
type UserPresenter interface {
	// PresentUser ユーザー情報を表示用の形式に変換する
	// @param user ユーザー情報のDTOオブジェクト
	// @return UserResponse 表示用のユーザー情報
	PresentUser(user *dto.UserResponse) UserResponse

	// PresentError エラー情報を表示用の形式に変換する
	// @param err エラーオブジェクト
	// @return ErrorResponse 表示用のエラー情報
	PresentError(err error) response.ErrorResponse
}

// UserResponse ユーザー情報のレスポンス構造体
// @Description ユーザー情報のレスポンス形式を定義する
// .
type UserResponse struct {
	// Status レスポンスのステータス
	Status string `example:"success" json:"status"`
	// Data ユーザー情報
	Data *dto.UserResponse `json:"data"`
}

type userPresenter struct{}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (p *userPresenter) PresentUser(user *dto.UserResponse) UserResponse {
	return UserResponse{
		Status: "success",
		Data:   user,
	}
}

func (p *userPresenter) PresentError(err error) response.ErrorResponse {
	return response.NewErrorResponse(err)
}
