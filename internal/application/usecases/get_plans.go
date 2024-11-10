package usecases

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
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

func (uc *GetPlansUseCase) Execute(userID string) ([]*dto.PlanResponse, error) {
	plans, err := uc.planRepo.FindAll(userID)
	if err != nil {
		return nil, err
	}

	response := make([]*dto.PlanResponse, len(plans))
	for i, plan := range plans {
		response[i] = dto.ToPlanResponse(plan)
	}

	return response, nil
}
