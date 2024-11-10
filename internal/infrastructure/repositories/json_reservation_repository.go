package repositories

import (
	"encoding/json"
	"fmt"
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

// 具象型のコンストラクタ.
func NewJSONReservationRepository(filePath string) repositories.ReservationRepository {
	return &jsonReservationRepository{
		filePath: filePath,
	}
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

// private メソッドとして findAll を移動.
func (r *jsonReservationRepository) findAll() ([]*models.Reservation, error) {
	absPath, err := filepath.Abs(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("ファイルパスの取得に失敗: %w", err)
	}

	file, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("ファイルの読み込みに失敗: %w", err)
	}

	var data reservationsJSON
	if err = json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("予約の取得に失敗: %w", err)
	}

	return data.Reservations, nil
}
