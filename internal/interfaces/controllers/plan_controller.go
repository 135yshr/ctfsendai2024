package controllers

import (
	"net/http"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/gin-gonic/gin"
)

type PlanController struct {
	getPlansUseCase *usecases.GetPlansUseCase
	presenter       presenters.PlanPresenter
}

func NewPlanController(
	getPlansUseCase *usecases.GetPlansUseCase,
	presenter presenters.PlanPresenter,
) *PlanController {
	return &PlanController{
		getPlansUseCase: getPlansUseCase,
		presenter:       presenter,
	}
}

// GetPlans ユーザーに関連するプラン一覧を取得します
// @Summary プラン一覧取得
// @Description 指定されたユーザーIDに関連するプラン一覧を取得します
// @Tags plans
// @Accept json
// @Produce json
// @Security Bearer
// @Param X-User-ID header string true "ユーザーID"
// @Success 200 {object} presenters.PlansResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /plans [get]
// .
func (c *PlanController) GetPlans(ctx *gin.Context) {
	userID := ctx.GetHeader("X-User-ID")

	plans, err := c.getPlansUseCase.Execute(userID)
	if err != nil {
		response := c.presenter.PresentError(err)
		ctx.JSON(http.StatusInternalServerError, response)

		return
	}

	response := c.presenter.PresentPlans(plans)
	ctx.JSON(http.StatusOK, response)
}
