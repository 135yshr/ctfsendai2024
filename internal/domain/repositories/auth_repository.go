package repositories

import (
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
)

type AuthRepository interface {
	FindByUserID(userID string) (*models.Auth, error)
	Store(auth *models.Auth) error
	ValidateToken(token string) (*models.Auth, error)
	GenerateToken(auth *models.Auth) (*models.Token, error)
}
