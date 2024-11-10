package repositories

import "github.com/135yshr/ctfsendai2024/internal/domain/models"

type PlanRepository interface {
	FindAll(userID string) ([]*models.Plan, error)
}
