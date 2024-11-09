package repositories

import "github.com/135yshr/ctfsendai2024/internal/domain/models"

type ReservationRepository interface {
	FindByUserID(userID string) ([]*models.Reservation, error)
}
