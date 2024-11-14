package usecases

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type SecretLoginUseCase struct {
	userRepository repositories.UserRepository
	authRepository repositories.AuthRepository
}

func NewSecretLoginUseCase(
	userRepository repositories.UserRepository,
	authRepository repositories.AuthRepository,
) *SecretLoginUseCase {
	return &SecretLoginUseCase{
		userRepository: userRepository,
		authRepository: authRepository,
	}
}

func (u *SecretLoginUseCase) Execute(
	ctx context.Context,
	userID string,
	secretAnswer string,
) (*dto.LoginResponse, error) {
	user, err := u.userRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("ユーザーが見つかりません: %w", err)
	}

	slog.InfoContext(ctx, "secret_answer", "secret_answer", secretAnswer, "user_secret_answer", user.SecretAnswer)
	if user.SecretAnswer != secretAnswer {
		return nil, errors.ErrInvalidSecretAnswer
	}

	auth := &models.Auth{
		UserID: user.ID,
		Name:   user.Name,
		Role:   user.Role,
	}
	token, err := u.authRepository.GenerateToken(ctx, auth)
	if err != nil {
		return nil, fmt.Errorf("トークン生成エラー: %w", err)
	}

	return &dto.LoginResponse{
		AccessToken: token.AccessToken,
		ExpiresAt:   token.ExpiresAt.Unix(),
	}, nil
}
