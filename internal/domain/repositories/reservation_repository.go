package repositories

import (
	"context"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
)

type ReservationRepository interface {
	FindByUserID(ctx context.Context, userID string) ([]*models.Reservation, error)
}
