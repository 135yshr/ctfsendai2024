package controllers

import (
	"net/http"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/validators"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
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
// @Param        user_id    query     string  true  "ユーザーID"  minlength(3)  maxlength(50)
// @Success      200  {object}  presenters.ReservationsResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /reservations [get]
func (c *ReservationController) GetUserReservations(ctx *gin.Context) {
	var req validators.GetReservationsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response := c.presenter.PresentError(err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	reservations, err := c.getUserReservationsUseCase.Execute(req.UserID)
	if err != nil {
		response := c.presenter.PresentError(err)
		ctx.JSON(http.StatusInternalServerError, response)

		return
	}

	response := c.presenter.PresentReservations(reservations)
	ctx.JSON(http.StatusOK, response)
}
