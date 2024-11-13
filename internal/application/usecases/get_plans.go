package usecases

import (
	"context"
	"fmt"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type GetPlansUseCase struct {
	planRepo repositories.PlanRepository
}

func NewGetPlansUseCase(
	repo repositories.PlanRepository,
) *GetPlansUseCase {
	return &GetPlansUseCase{
		planRepo: repo,
	}
}

func (uc *GetPlansUseCase) Execute(ctx context.Context, params *models.PlanSearchParams) ([]*dto.PlanResponse, error) {
	plans, err := uc.planRepo.FindAll(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("プランの取得に失敗: %w", err)
	}

	response := make([]*dto.PlanResponse, len(plans))
	for i, plan := range plans {
		response[i] = dto.ToPlanResponse(plan)
	}

	return response, nil
}
