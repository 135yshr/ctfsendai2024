package controllers

import (
	"errors"
	"net/http"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	domainError "github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	loginUseCase *usecases.LoginUseCase
	presenter    presenters.AuthPresenter
}

func NewAuthController(
	loginUseCase *usecases.LoginUseCase,
	presenter presenters.AuthPresenter,
) *AuthController {
	return &AuthController{
		loginUseCase: loginUseCase,
		presenter:    presenter,
	}
}

// @Summary ユーザーログイン
// @Description ユーザー名とパスワードを使用してログイン認証を行います
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "ログイン情報"
// @Success 200 {object} presenters.LoginResponse "ログイン成功時のレスポンス"
// @Failure 400 {object} presenters.PresentError "リクエストの形式が不正"
// @Failure 401 {object} presenters.PresentError "パスワードが一致しない"
// @Failure 500 {object} presenters.PresentError "サーバー内部エラー"
// @Router /login [post].
func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		response := c.presenter.PresentError(err)
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	result, err := c.loginUseCase.Execute(
		ctx,
		loginRequest.UserID,
		loginRequest.Password,
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
