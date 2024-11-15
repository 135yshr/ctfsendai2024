package controllers

import (
	"net/http"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	domainError "github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/middleware"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/validators"
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
// @Summary プラン一覧取得API
// @Description ユーザーの権限に基づいて、指定された検索条件に一致するプラン一覧を取得します
// @Tags plans
// @Accept json
// @Produce json
// @Security Bearer
// @Param startDate query string false "開始日 (YYYY-MM-DD形式)"
// @Param endDate query string false "終了日 (YYYY-MM-DD形式)"
// @Param status query string false "プランのステータス (active/inactive/pending)"
// @Success 200 {object} presenters.PlansResponse "プラン一覧の取得に成功"
// @Failure 400 {object} response.ErrorResponse "不正なリクエストパラメータ"
// @Failure 401 {object} response.ErrorResponse "認証エラー"
// @Failure 500 {object} response.ErrorResponse "サーバー内部エラー"
// @Router /plans [get]
// .
func (pc *PlanController) GetPlans(c *gin.Context) {
	auth, ok := utils.GetUserFromContext(c)
	if !ok {
		response := pc.presenter.PresentError(domainError.ErrInvalidUser)
		c.JSON(http.StatusUnauthorized, response)

		return
	}

	var req validators.GetPlansRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response := pc.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	params := &models.PlanSearchParams{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Status:    &req.Status,
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
