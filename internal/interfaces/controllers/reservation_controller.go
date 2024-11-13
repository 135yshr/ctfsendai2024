package controllers

import (
	"net/http"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	domainError "github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/middleware"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/validators"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/utils"
	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	getUserReservationsUseCase *usecases.GetUserReservationsUseCase
	presenter                  presenters.ReservationPresenter
}

func NewReservationController(
	getUserReservationsUseCase *usecases.GetUserReservationsUseCase,
	presenter presenters.ReservationPresenter,
) *ReservationController {
	return &ReservationController{
		getUserReservationsUseCase: getUserReservationsUseCase,
		presenter:                  presenter,
	}
}

// GetUserReservations ユーザーの予約一覧を取得
// @Summary      ユーザーの予約一覧を取得
// @Description  指定されたユーザーIDに紐づく予約の一覧を取得します
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Security Bearer
// @Param        user_id    query     string  true  "ユーザーID"  minlength(3)  maxlength(50)
// @Success      200  {object}  presenters.ReservationsResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /reservations [get]
// .
func (rc *ReservationController) GetUserReservations(c *gin.Context) {
	auth, ok := utils.GetUserFromContext(c)
	if !ok {
		response := rc.presenter.PresentError(domainError.ErrInvalidUser)
		c.JSON(http.StatusUnauthorized, response)

		return
	}

	var req validators.GetReservationsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response := rc.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	ctx := middleware.SetUserToContext(c, auth)
	reservations, err := rc.getUserReservationsUseCase.Execute(ctx, req.UserID)
	if err != nil {
		response := rc.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	response := rc.presenter.PresentReservations(reservations)
	c.JSON(http.StatusOK, response)
}
