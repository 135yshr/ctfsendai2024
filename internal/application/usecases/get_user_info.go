package usecases

import (
	"context"
	"fmt"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type GetUserInfoUseCase struct {
	userRepo repositories.UserRepository
}

func NewGetUserInfoUseCase(userRepo repositories.UserRepository) *GetUserInfoUseCase {
	return &GetUserInfoUseCase{
		userRepo: userRepo,
	}
}

func (u *GetUserInfoUseCase) Execute(ctx context.Context, userID string) (*dto.UserResponse, error) {
	user, err := u.userRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("ユーザー情報の取得に失敗しました: %w", err)
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}
