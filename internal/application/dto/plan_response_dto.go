package dto

import "github.com/135yshr/ctfsendai2024/internal/domain/models"

// PlanResponse プラン情報レスポンス.
type PlanResponse struct {
	// プランID
	ID string `example:"plan123" json:"id"`

	// プラン名
	Name string `example:"スタンダードプラン" json:"name"`

	// プランの説明
	Description string `example:"基本的なサービスが含まれるプランです" json:"description"`

	// プランの価格
	Price int `example:"1000" json:"price"`

	// プランの期間（日数）
	Duration int `example:"30" json:"duration"`
}

func ToPlanResponse(plan *models.Plan) *PlanResponse {
	return &PlanResponse{
		ID:          plan.ID,
		Name:        plan.Name,
		Description: plan.Description,
		Price:       plan.Price,
		Duration:    plan.Duration,
	}
}
