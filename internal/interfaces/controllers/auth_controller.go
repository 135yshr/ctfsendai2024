package controllers

import (
	"errors"
	"net/http"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	domainError "github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/validators"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	loginUseCase             *usecases.LoginUseCase
	secretLoginUseCase       *usecases.SecretLoginUseCase
	getSecretQuestionUseCase *usecases.GetSecretQuestionUseCase
	presenter                presenters.AuthPresenter
}

func NewAuthController(
	loginUseCase *usecases.LoginUseCase,
	secretLoginUseCase *usecases.SecretLoginUseCase,
	getSecretQuestionUseCase *usecases.GetSecretQuestionUseCase,
	presenter presenters.AuthPresenter,
) *AuthController {
	return &AuthController{
		loginUseCase:             loginUseCase,
		secretLoginUseCase:       secretLoginUseCase,
		getSecretQuestionUseCase: getSecretQuestionUseCase,
		presenter:                presenter,
	}
}

// @Summary ユーザーログイン
// @Description ユーザー名とパスワードを使用してログイン認証を行います
// @Tags auth
// @Accept json
// @Produce json
// @Param request body validators.LoginRequest true "ログイン情報"
// @Success 200 {object} presenters.LoginResponse "ログイン成功時のレスポンス"
// @Failure 400 {object} presenters.PresentError "リクエストの形式が不正"
// @Failure 401 {object} presenters.PresentError "パスワードが一致しない"
// @Failure 500 {object} presenters.PresentError "サーバー内部エラー"
// @Router /login [post].
func (c *AuthController) Login(ctx *gin.Context) {
	var req validators.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := c.presenter.PresentError(err)
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	result, err := c.loginUseCase.Execute(
		ctx,
		req.UserID,
		req.Password,
	)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, domainError.ErrInvalidPassword) {
			statusCode = http.StatusUnauthorized
		}
		response := c.presenter.PresentError(err)
		ctx.JSON(statusCode, response)

		return
	}

	response := c.presenter.PresentLogin(result)
	ctx.JSON(http.StatusOK, response)
}

// @Summary 秘密の質問によるログイン
// @Description 秘密の質問の回答を使用してログイン認証を行います
// @Tags auth
// @Accept json
// @Produce json
// @Param request body validators.SecretLoginRequest true "秘密の質問の回答情報"
// @Success 200 {object} presenters.LoginResponse "ログイン成功時のレスポンス"
// @Failure 400 {object} presenters.PresentError "リクエストの形式が不正"
// @Failure 401 {object} presenters.PresentError "秘密の質問の回答が一致しない"
// @Failure 500 {object} presenters.PresentError "サーバー内部エラー"
// @Router /secret-login [post]
// .
func (c *AuthController) SecretLogin(ctx *gin.Context) {
	var req validators.SecretLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := c.presenter.PresentError(err)
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	result, err := c.secretLoginUseCase.Execute(
		ctx,
		req.UserID,
		req.SecretAnswer,
	)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, domainError.ErrInvalidSecretAnswer) {
			statusCode = http.StatusUnauthorized
		}
		response := c.presenter.PresentError(err)
		ctx.JSON(statusCode, response)

		return
	}

	response := c.presenter.PresentLogin(result)
	ctx.JSON(http.StatusOK, response)
}

// @Summary 秘密の質問の取得
// @Description ユーザーIDに対応する秘密の質問を取得します
// @Tags auth
// @Accept json
// @Produce json
// @Param user_id query string true "ユーザーID"
// @Success 200 {object} dto.SecretQuestionResponse "秘密の質問"
// @Failure 400 {object} presenters.PresentError "リクエストの形式が不正"
// @Failure 404 {object} presenters.PresentError "ユーザーが見つからない"
// @Failure 500 {object} presenters.PresentError "サーバー内部エラー"
// @Router /secret-question [get]
// .
func (c *AuthController) GetSecretQuestion(ctx *gin.Context) {
	var req validators.SecretQuestionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response := c.presenter.PresentError(err)
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	result, err := c.getSecretQuestionUseCase.Execute(ctx, req.UserID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, domainError.ErrUserNotFound) {
			statusCode = http.StatusNotFound
		}
		response := c.presenter.PresentError(err)
		ctx.JSON(statusCode, response)

		return
	}

	response := c.presenter.PresentSecretQuestion(result)
	ctx.JSON(http.StatusOK, response)
}
