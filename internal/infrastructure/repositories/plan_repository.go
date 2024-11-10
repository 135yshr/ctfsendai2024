package repositories

import (
	"encoding/json"
	"os"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type jsonPlanRepository struct {
	dbPath string
}

func NewJSONPlanRepository(impl *jsonPlanRepository) repositories.PlanRepository {
	return impl
}

func NewJSONPlanRepositoryImpl(dbPath string) *jsonPlanRepository {
	return &jsonPlanRepository{
		dbPath: dbPath,
	}
}

func (r *jsonPlanRepository) FindAll(userID string) ([]*models.Plan, error) {
	file, err := os.ReadFile(r.dbPath)
	if err != nil {
		return nil, err
	}

	var data struct {
		Plans []*models.Plan `json:"plans"`
	}
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	var plans []*models.Plan
	for _, plan := range data.Plans {
		if plan.ID == "p000" {
			if userID == "u00000" {
				plans = append(plans, plan)
			}
			continue
		}
		plans = append(plans, plan)
	}

	return plans, nil
}
