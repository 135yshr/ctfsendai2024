package usecases

import (
	"context"
	"fmt"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type GetUsersUseCase struct {
	userRepository repositories.UserRepository
}

func NewGetUsersUseCase(userRepository repositories.UserRepository) *GetUsersUseCase {
	return &GetUsersUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUsersUseCase) Execute(ctx context.Context) ([]*dto.UserResponse, error) {
	users, err := uc.userRepository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	userResponses := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = &dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		}
	}

	return userResponses, nil
}
