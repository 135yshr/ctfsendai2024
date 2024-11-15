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

	// PresentUsers ユーザー情報を表示用の形式に変換する
	// @param users ユーザー情報のDTOオブジェクト
	// @return UsersResponse 表示用のユーザー情報
	PresentUsers(users []*dto.UserResponse) UsersResponse

	// PresentError エラー情報を表示用の形式に変換する
	// @param err エラーオブジェクト
	// @return ErrorResponse 表示用のエラー情報
	PresentError(err error) response.ErrorResponse
}

type userPresenter struct{}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
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

func (p *userPresenter) PresentUser(user *dto.UserResponse) UserResponse {
	return UserResponse{
		Status: "success",
		Data:   user,
	}
}

// UsersResponse ユーザー情報のレスポンス構造体
// @Description ユーザー情報のレスポンス形式を定義する
// .
type UsersResponse struct {
	Status string              `example:"success" json:"status"`
	Data   []*dto.UserResponse `json:"data"`
}

func (p *userPresenter) PresentUsers(users []*dto.UserResponse) UsersResponse {
	return UsersResponse{
		Status: "success",
		Data:   users,
	}
}

func (p *userPresenter) PresentError(err error) response.ErrorResponse {
	return response.NewErrorResponse(err)
}
