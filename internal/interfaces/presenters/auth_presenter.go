package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters/response"
)

type AuthPresenter interface {
	PresentError(err error) response.ErrorResponse
	PresentLogin(data *dto.LoginResponse) LoginResponse
	PresentSecretQuestion(question *dto.SecretQuestionResponse) SecretQuestionResponse
}

type authPresenter struct{}

func NewAuthPresenter() AuthPresenter {
	return &authPresenter{}
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

func (p *authPresenter) PresentLogin(data *dto.LoginResponse) LoginResponse {
	return LoginResponse{
		Status: "success",
		Data:   data,
	}
}

type SecretQuestionResponse struct {
	SecretQuestion string `json:"secret_question"`
}

func (p *authPresenter) PresentSecretQuestion(question *dto.SecretQuestionResponse) SecretQuestionResponse {
	return SecretQuestionResponse{
		SecretQuestion: question.SecretQuestion,
	}
}

func (p *authPresenter) PresentError(err error) response.ErrorResponse {
	return response.NewErrorResponse(err)
}
