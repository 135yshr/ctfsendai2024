package repositories

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type jsonReservationRepository struct {
	filePath string
}

type reservationsJSON struct {
	Reservations []*models.Reservation `json:"reservations"`
}

// 具象型のコンストラクタ
func NewJSONReservationRepositoryImpl(filePath string) *jsonReservationRepository {
	return &jsonReservationRepository{
		filePath: filePath,
	}
}

// インターフェースを返すコンストラクタ
func NewJSONReservationRepository(impl *jsonReservationRepository) repositories.ReservationRepository {
	return impl
}

func (r *jsonReservationRepository) FindByUserID(userID string) ([]*models.Reservation, error) {
	reservations, err := r.findAll()
	if err != nil {
		return nil, err
	}

	var result []*models.Reservation
	for _, reservation := range reservations {
		if reservation.UserID == userID {
			result = append(result, reservation)
		}
	}

	return result, nil
}

// private メソッドとして findAll を移動
func (r *jsonReservationRepository) findAll() ([]*models.Reservation, error) {
	absPath, err := filepath.Abs(r.filePath)
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	var data reservationsJSON
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data.Reservations, nil
}
