package usecases

import (
	"context"
	"fmt"

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

func (uc *GetUserReservationsUseCase) Execute(ctx context.Context, userID string) ([]*dto.ReservationResponse, error) {
	// リポジトリから予約一覧を取得
	reservations, err := uc.reservationRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("予約の取得に失敗: %w", err)
	}

	// ドメインモデルをDTOに変換
	response := make([]*dto.ReservationResponse, len(reservations))
	for i, reservation := range reservations {
		response[i] = dto.ToReservationResponse(reservation)
	}

	return response, nil
}
