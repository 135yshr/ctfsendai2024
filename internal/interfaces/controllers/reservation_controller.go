package controllers

import (
	"net/http"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	domainError "github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/middleware"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/validators"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/utils"
	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	presenter                  presenters.ReservationPresenter
	getUserReservationsUseCase *usecases.GetUserReservationsUseCase
	createReservationUseCase   *usecases.CreateReservationUseCase
}

func NewReservationController(
	getUserReservationsUseCase *usecases.GetUserReservationsUseCase,
	createReservationUseCase *usecases.CreateReservationUseCase,
	presenter presenters.ReservationPresenter,
) *ReservationController {
	return &ReservationController{
		getUserReservationsUseCase: getUserReservationsUseCase,
		createReservationUseCase:   createReservationUseCase,
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

// CreateReservation 新しい予約を作成
// @Summary      新しい予約を作成
// @Description  ユーザーの新しい予約を作成します
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        request body validators.CreateReservationRequest true "予約情報"
// @Success      201  {object}  presenters.ReservationResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      401  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /reservations [post]
// .
func (rc *ReservationController) CreateReservation(c *gin.Context) {
	auth, ok := utils.GetUserFromContext(c)
	if !ok {
		response := rc.presenter.PresentError(domainError.ErrInvalidUser)
		c.JSON(http.StatusUnauthorized, response)

		return
	}

	var req validators.CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response := rc.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	ctx := middleware.SetUserToContext(c, auth)
	reservation, err := rc.createReservationUseCase.Execute(ctx, &dto.CreateReservationRequest{
		UserID:    req.UserID,
		PlanID:    req.PlanID,
		StartDate: req.StartDate,
	})
	if err != nil {
		response := rc.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	response := rc.presenter.PresentReservation(reservation)
	c.JSON(http.StatusCreated, response)
}
