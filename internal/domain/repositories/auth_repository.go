package repositories

import (
	"context"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
)

type AuthRepository interface {
	ValidateToken(ctx context.Context, token string) (*models.Auth, error)
	GenerateToken(ctx context.Context, auth *models.Auth) (*models.Token, error)
}
