package usecases

import (
	"fmt"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	authRepository repositories.AuthRepository
}

func NewLoginUseCase(authRepo repositories.AuthRepository) *LoginUseCase {
	return &LoginUseCase{
		authRepository: authRepo,
	}
}

func (uc *LoginUseCase) Execute(username, password string) (*dto.LoginResponse, error) {
	// ユーザー認証
	auth, err := uc.authRepository.FindByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("ユーザーが見つかりません: %w", err)
	}

	// パスワードの検証
	if !comparePasswords(auth.Password, password) {
		return nil, errors.ErrInvalidPassword
	}

	// トークンの生成
	token, err := uc.authRepository.GenerateToken(auth)
	if err != nil {
		return nil, fmt.Errorf("トークン生成エラー: %w", err)
	}

	// レスポンスの作成
	response := &dto.LoginResponse{
		AccessToken: token.AccessToken,
		ExpiresAt:   token.ExpiresAt.Unix(),
	}

	return response, nil
}

// パスワード比較のヘルパー関数.
func comparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	return err == nil
}
