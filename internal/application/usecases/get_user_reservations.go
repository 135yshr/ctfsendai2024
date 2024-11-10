package usecases

import (
	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type GetUserReservationsUseCase struct {
	reservationRepo repositories.ReservationRepository
}

func NewGetUserReservationsUseCase(
	repo repositories.ReservationRepository,
) *GetUserReservationsUseCase {
	return &GetUserReservationsUseCase{
		reservationRepo: repo,
	}
}

func (uc *GetUserReservationsUseCase) Execute(userID string) ([]*dto.ReservationResponse, error) {
	// リポジトリから予約一覧を取得
	reservations, err := uc.reservationRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	// ドメインモデルをDTOに変換
	response := make([]*dto.ReservationResponse, len(reservations))
	for i, reservation := range reservations {
		response[i] = dto.ToReservationResponse(reservation)
	}

	return response, nil
}
