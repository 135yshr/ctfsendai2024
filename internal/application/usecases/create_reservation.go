package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type CreateReservationUseCase struct {
	reservationRepo repositories.ReservationRepository
	userRepo        repositories.UserRepository
	planRepo        repositories.PlanRepository
}

func NewCreateReservationUseCase(
	reservationRepo repositories.ReservationRepository,
	userRepo repositories.UserRepository,
	planRepo repositories.PlanRepository,
) *CreateReservationUseCase {
	return &CreateReservationUseCase{
		reservationRepo: reservationRepo,
		userRepo:        userRepo,
		planRepo:        planRepo,
	}
}

func (u *CreateReservationUseCase) Execute(
	ctx context.Context,
	req *dto.CreateReservationRequest,
) (*dto.ReservationResponse, error) {
	user, err := u.userRepo.FindByUserID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("ユーザーが見つかりません: %w", err)
	}

	plan, err := u.planRepo.FindByID(ctx, req.PlanID)
	if err != nil {
		return nil, fmt.Errorf("プランが見つかりません: %w", err)
	}

	// 予約を作成
	reservation := &models.Reservation{
		UserID:    user.ID,
		PlanID:    req.PlanID,
		StartTime: req.StartDate,
		EndTime:   req.StartDate.Add(time.Duration(plan.Duration) * time.Minute),
		Status:    models.StatusReserved,
		User:      user,
		Plan:      plan,
	}
	reservation, err = u.reservationRepo.Create(ctx, reservation)
	if err != nil {
		return nil, fmt.Errorf("予約の作成に失敗しました: %w", err)
	}

	return dto.ToReservationResponse(reservation), nil
}
