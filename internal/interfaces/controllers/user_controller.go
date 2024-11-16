package controllers

import (
	"net/http"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	"github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	presenter          presenters.UserPresenter
	getUserInfoUseCase *usecases.GetUserInfoUseCase
	getUsersUseCase    *usecases.GetUsersUseCase
}

func NewUserController(
	presenter presenters.UserPresenter,
	getUserInfoUseCase *usecases.GetUserInfoUseCase,
	getUsersUseCase *usecases.GetUsersUseCase,
) *UserController {
	return &UserController{
		presenter:          presenter,
		getUserInfoUseCase: getUserInfoUseCase,
		getUsersUseCase:    getUsersUseCase,
	}
}

// GetMe ログイン中のユーザー情報を取得します
//
//	@Summary		ログインユーザー情報取得
//	@Description	ログイン中のユーザー情報を取得します
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Success		200	{object}	presenters.UserResponse
//	@Failure		401	{object}	response.ErrorResponse
//	@Router			/users/me [get]
//
// .
func (uc *UserController) GetMe(c *gin.Context) {
	auth, ok := utils.GetUserFromContext(c)
	if !ok {
		response := uc.presenter.PresentError(errors.ErrInvalidUser)
		c.JSON(http.StatusUnauthorized, response)

		return
	}

	result, err := uc.getUserInfoUseCase.Execute(c, auth.UserID)
	if err != nil {
		response := uc.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	response := uc.presenter.PresentUser(result)
	c.JSON(http.StatusOK, response)
}

// GetUsers 全ユーザー情報を取得します
//
//	@Summary		ユーザー一覧取得
//	@Description	全ユーザーの情報を取得します（管理者のみ）
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Success		200	{array}		presenters.UserResponse
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Router			/users [get]
//
// .
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.getUsersUseCase.Execute(c)
	if err != nil {
		response := uc.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	response := uc.presenter.PresentUsers(users)
	c.JSON(http.StatusOK, response)
}
