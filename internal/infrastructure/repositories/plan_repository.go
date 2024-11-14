package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/middleware"
)

type planRepository struct {
	dbPath string
	plans  []*models.Plan
}

func NewPlanRepository(dbPath string) (repositories.PlanRepository, error) {
	repo := &planRepository{
		dbPath: dbPath,
	}

	// 初期化時にファイルを読み込む
	if err := repo.loadPlans(); err != nil {
		return nil, fmt.Errorf("プランの初期化に失敗: %w", err)
	}

	return repo, nil
}

func (r *planRepository) loadPlans() error {
	file, err := os.ReadFile(r.dbPath)
	if err != nil {
		return fmt.Errorf("ファイルの読み込みに失敗: %w", err)
	}

	var data struct {
		Plans []*models.Plan `json:"plans"`
	}
	if err = json.Unmarshal(file, &data); err != nil {
		return fmt.Errorf("プランの取得に失敗: %w", err)
	}

	r.plans = data.Plans

	return nil
}

func (r *planRepository) FindAll(ctx context.Context, _ *models.PlanSearchParams) ([]*models.Plan, error) {
	user, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ユーザーの取得に失敗: %w", err)
	}

	plans := []*models.Plan{}
	for _, plan := range r.plans {
		if plan.ID == "p000" {
			if user.Role == models.RoleAdmin {
				plans = append(plans, plan)
			}

			continue
		}
		plans = append(plans, plan)
	}

	return plans, nil
}

func (r *planRepository) FindByID(_ context.Context, id string) (*models.Plan, error) {
	for _, plan := range r.plans {
		if plan.ID == id {
			return plan, nil
		}
	}

	return nil, errors.ErrPlanNotFound
}
