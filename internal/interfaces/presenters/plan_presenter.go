package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters/response"
)

type PlanPresenter interface {
	PresentPlans(plans []*dto.PlanResponse) PlansResponse
	PresentError(err error) response.ErrorResponse
}

type JSONPlanPresenter struct{}

func NewJSONPlanPresenter() PlanPresenter {
	return &JSONPlanPresenter{}
}

// PlansResponse プラン一覧レスポンス
// @Description プラン一覧のレスポンス.
type PlansResponse struct {
	// ステータス
	// @Example "success"
	Status string `example:"success" json:"status"`

	// プランデータ
	Data []*dto.PlanResponse `json:"data"`
}

func (p *JSONPlanPresenter) PresentPlans(plans []*dto.PlanResponse) PlansResponse {
	return PlansResponse{
		Status: "success",
		Data:   plans,
	}
}

func (p *JSONPlanPresenter) PresentError(err error) response.ErrorResponse {
	return response.NewErrorResponse(err)
}
