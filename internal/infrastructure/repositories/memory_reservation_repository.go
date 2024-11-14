package repositories

import (
	"context"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type memoryReservationRepository struct {
	reservations []*models.Reservation
}

func NewMemoryReservationRepository() repositories.ReservationRepository {
	return &memoryReservationRepository{
		reservations: make([]*models.Reservation, 0),
	}
}

func (r *memoryReservationRepository) FindByUserID(_ context.Context, userID string) ([]*models.Reservation, error) {
	var result []*models.Reservation
	for _, reservation := range r.reservations {
		if reservation.UserID == userID {
			result = append(result, reservation)
		}
	}

	return result, nil
}

func (r *memoryReservationRepository) Create(
	_ context.Context,
	reservation *models.Reservation,
) (*models.Reservation, error) {
	r.reservations = append(r.reservations, reservation)

	return reservation, nil
}
