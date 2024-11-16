package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters/response"
)

// UserPresenter ユーザー情報のプレゼンテーション層のインターフェース
//
//	@Description	ユーザー情報の表示形式を管理するプレゼンター.
type UserPresenter interface {
	// PresentUser ユーザー情報を表示用の形式に変換する
	//	@Description	単一のユーザー情報を成功レスポンスの形式に変換する
	//	@Param			user	変換対象のユーザー情報
	//	@Return			UserResponse 変換されたユーザーレスポンス
	PresentUser(user *dto.UserResponse) UserResponse

	// PresentUsers ユーザー情報を表示用の形式に変換する
	//	@Description	複数のユーザー情報を成功レスポンスの形式に変換する
	//	@Param			users	変換対象のユーザー情報一覧
	//	@Return			UsersResponse 変換されたユーザー一覧レスポンス
	PresentUsers(users []*dto.UserResponse) UsersResponse

	// PresentError エラー情報を表示用の形式に変換する
	//	@Description	エラー情報をエラーレスポンスの形式に変換する
	//	@Param			err	変換対象のエラー
	//	@Return			ErrorResponse 変換されたエラーレスポンス
	PresentError(err error) response.ErrorResponse
}

type userPresenter struct{}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

// UserResponse ユーザー情報のレスポンス構造体
//
//	@Description	ユーザー情報の単一レスポンス形式を定義する.
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
//
//	@Description	ユーザー情報の一覧レスポンス形式を定義する.
type UsersResponse struct {
	// Status レスポンスのステータス
	Status string `example:"success" json:"status"`

	// Data ユーザー情報一覧
	Data []*dto.UserResponse `json:"data"`
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
