package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"github.com/google/uuid"
)

type jsonReservationRepository struct {
	filePath     string
	reservations map[string]*models.Reservation
}

type reservationsJSON struct {
	Reservations []*models.Reservation `json:"reservations"`
}

// 具象型のコンストラクタ.
func NewJSONReservationRepository(filePath string) (repositories.ReservationRepository, error) {
	repo := &jsonReservationRepository{
		filePath:     filePath,
		reservations: make(map[string]*models.Reservation),
	}

	// 初期化時にファイルを読み込む
	reservations, err := repo.loadReservations()
	if err != nil {
		return nil, fmt.Errorf("予約データの初期化に失敗: %w", err)
	}

	for _, reservation := range reservations {
		repo.reservations[reservation.ID] = reservation
	}

	return repo, nil
}

func (r *jsonReservationRepository) FindByUserID(_ context.Context, userID string) ([]*models.Reservation, error) {
	var reservations []*models.Reservation
	for _, v := range r.reservations {
		if v.UserID == userID {
			reservations = append(reservations, v)
		}
	}

	return reservations, nil
}

// private メソッドとして loadReservations を移動.
func (r *jsonReservationRepository) loadReservations() ([]*models.Reservation, error) {
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

// Create は予約を作成します.
func (r *jsonReservationRepository) Create(
	_ context.Context,
	reservation *models.Reservation,
) (*models.Reservation, error) {
	reservation.ID = uuid.New().String()
	r.reservations[reservation.ID] = reservation

	return reservation, nil
}
