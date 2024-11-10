package repositories

import (
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

func (r *memoryReservationRepository) FindByUserID(userID string) ([]*models.Reservation, error) {
	var result []*models.Reservation
	for _, reservation := range r.reservations {
		if reservation.UserID == userID {
			result = append(result, reservation)
		}
	}

	return result, nil
}