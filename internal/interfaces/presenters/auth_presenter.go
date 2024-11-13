package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
)

// PresentError エラーレスポンスの構造体.
type PresentError struct {
	Message string `json:"message"`
}

type AuthPresenter interface {
	PresentError(err error) PresentError
	PresentLogin(data *dto.LoginResponse) LoginResponse
}

type JSONAuthPresenter struct{}

func NewJSONAuthPresenter() AuthPresenter {
	return &JSONAuthPresenter{}
}

// LoginResponse ログインレスポンス
// @Description ログイン処理のレスポンス.
type LoginResponse struct {
	// ステータス
	// @Example "success"
	Status string `example:"success" json:"status"`

	// ログインデータ
	Data *dto.LoginResponse `json:"data"`
}

func (p *JSONAuthPresenter) PresentLogin(data *dto.LoginResponse) LoginResponse {
	return LoginResponse{
		Status: "success",
		Data:   data,
	}
}

func (p *JSONAuthPresenter) PresentError(err error) PresentError {
	return PresentError{
		Message: err.Error(),
	}
}
