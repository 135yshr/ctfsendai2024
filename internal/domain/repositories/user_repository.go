package repositories

import (
	"context"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
)

type UserRepository interface {
	FindByUserID(ctx context.Context, userID string) (*models.User, error)
	Store(ctx context.Context, user *models.User) error
}
