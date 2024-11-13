package controllers

import (
	"net/http"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	domainError "github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/middleware"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/utils"
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

// GetPlans プラン一覧を取得します
// @Summary プラン一覧取得
// @Description 検索条件に基づいてプラン一覧を取得します
// @Tags plans
// @Accept json
// @Produce json
// @Security Bearer
// @Param startDate query string false "開始日 (YYYY-MM-DD)"
// @Param endDate query string false "終了日 (YYYY-MM-DD)"
// @Param status query string false "ステータス"
// @Success 200 {object} presenters.PlansResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /plans [get]
// .
func (pc *PlanController) GetPlans(c *gin.Context) {
	auth, ok := utils.GetUserFromContext(c)
	if !ok {
		response := pc.presenter.PresentError(domainError.ErrInvalidUser)
		c.JSON(http.StatusUnauthorized, response)

		return
	}

	params := &models.PlanSearchParams{}
	if err := c.ShouldBind(params); err != nil {
		response := pc.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	ctx := middleware.SetUserToContext(c, auth)
	plans, err := pc.getPlansUseCase.Execute(ctx, params)
	if err != nil {
		response := pc.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	response := pc.presenter.PresentPlans(plans)
	c.JSON(http.StatusOK, response)
}
