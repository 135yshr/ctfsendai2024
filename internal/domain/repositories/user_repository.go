package repositories

import (
	"context"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]*models.User, error)
	FindByUserID(ctx context.Context, userID string) (*models.User, error)
	Store(ctx context.Context, user *models.User) error
}
