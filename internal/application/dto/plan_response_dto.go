package dto

import "github.com/135yshr/ctfsendai2024/internal/domain/models"

// PlanResponse プラン情報レスポンス
// @Description プラン情報の詳細
type PlanResponse struct {
	// プランID
	// @Example "plan123"
	ID string `json:"id" example:"plan123"`

	// プラン名
	// @Example "スタンダードプラン"
	Name string `json:"name" example:"スタンダードプラン"`

	// プランの説明
	// @Example "基本的なサービスが含まれるプランです"
	Description string `json:"description" example:"基本的なサービスが含まれるプランです"`

	// プランの価格
	// @Example 1000
	Price int `json:"price" example:"1000"`

	// プランの期間（日数）
	// @Example 30
	Duration int `json:"duration" example:"30"`
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
