package repositories

import (
	"context"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
)

type PlanRepository interface {
	FindAll(ctx context.Context, params *models.PlanSearchParams) ([]*models.Plan, error)
	FindByID(ctx context.Context, id string) (*models.Plan, error)
}
