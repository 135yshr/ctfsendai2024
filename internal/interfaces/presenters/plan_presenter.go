package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters/response"
)

type PlanPresenter interface {
	PresentPlans([]*dto.PlanResponse) interface{}
	PresentError(error) interface{}
}

type JSONPlanPresenter struct{}

func NewJSONPlanPresenter() PlanPresenter {
	return &JSONPlanPresenter{}
}

// PlansResponse プラン一覧レスポンス
// @Description プラン一覧のレスポンス
type PlansResponse struct {
	// ステータス
	// @Example "success"
	Status string `json:"status" example:"success"`

	// プランデータ
	Data []*dto.PlanResponse `json:"data"`
}

func (p *JSONPlanPresenter) PresentPlans(plans []*dto.PlanResponse) interface{} {
	return PlansResponse{
		Status: "success",
		Data:   plans,
	}
}

func (p *JSONPlanPresenter) PresentError(err error) interface{} {
	return response.NewErrorResponse(err)
}
