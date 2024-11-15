package presenters

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters/response"
)

// PlanPresenter プラン情報のプレゼンテーション層のインターフェース
// @Description プラン情報の表示形式を管理するプレゼンター.
type PlanPresenter interface {
	// PresentPlans プラン情報一覧を表示用の形式に変換する
	// @Description プラン情報一覧を成功レスポンスの形式に変換する
	// @Param plans 変換対象のプラン情報一覧
	// @Return PlansResponse 変換されたプラン一覧レスポンス
	PresentPlans(plans []*dto.PlanResponse) PlansResponse

	// PresentError エラー情報を表示用の形式に変換する
	// @Description エラー情報をエラーレスポンスの形式に変換する
	// @Param err 変換対象のエラー
	// @Return ErrorResponse 変換されたエラーレスポンス
	PresentError(err error) response.ErrorResponse
}

type planPresenter struct{}

func NewPlanPresenter() PlanPresenter {
	return &planPresenter{}
}

// PlansResponse プラン一覧レスポンス
// @Description プラン一覧のレスポンス形式を定義する.
type PlansResponse struct {
	// Status レスポンスのステータス
	// @Example "success"
	Status string `example:"success" json:"status"`

	// Data プラン情報一覧
	// @Description プラン情報の配列
	Data []*dto.PlanResponse `json:"data"`
}

func (p *planPresenter) PresentPlans(plans []*dto.PlanResponse) PlansResponse {
	return PlansResponse{
		Status: "success",
		Data:   plans,
	}
}

func (p *planPresenter) PresentError(err error) response.ErrorResponse {
	return response.NewErrorResponse(err)
}
