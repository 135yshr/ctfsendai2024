package repositories

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type jsonPlanRepository struct {
	dbPath string
}

func NewJSONPlanRepository(dbPath string) repositories.PlanRepository {
	return &jsonPlanRepository{
		dbPath: dbPath,
	}
}

func (r *jsonPlanRepository) FindAll(userID string) ([]*models.Plan, error) {
	file, err := os.ReadFile(r.dbPath)
	if err != nil {
		return nil, fmt.Errorf("ファイルの読み込みに失敗: %w", err)
	}

	var data struct {
		Plans []*models.Plan `json:"plans"`
	}
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("プランの取得に失敗: %w", err)
	}

	plans := []*models.Plan{}
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
