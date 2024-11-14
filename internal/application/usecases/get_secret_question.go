package usecases

import (
	"context"
	"fmt"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type GetSecretQuestionUseCase struct {
	userRepository repositories.UserRepository
}

func NewGetSecretQuestionUseCase(
	userRepository repositories.UserRepository,
) *GetSecretQuestionUseCase {
	return &GetSecretQuestionUseCase{
		userRepository: userRepository,
	}
}

func (u *GetSecretQuestionUseCase) Execute(
	ctx context.Context,
	userID string,
) (*dto.SecretQuestionResponse, error) {
	user, err := u.userRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("ユーザーが見つかりません: %w", err)
	}

	return &dto.SecretQuestionResponse{
		SecretQuestion: user.SecretQuestion,
	}, nil
}
