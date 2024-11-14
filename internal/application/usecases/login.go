package usecases

import (
	"context"
	"crypto/md5" //nolint:gosec // あえて脆弱なハッシュ関数を使用
	"encoding/hex"
	"fmt"

	"github.com/135yshr/ctfsendai2024/internal/application/dto"
	"github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
)

type LoginUseCase struct {
	authRepository repositories.AuthRepository
	userRepository repositories.UserRepository
}

func NewLoginUseCase(authRepo repositories.AuthRepository, userRepo repositories.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		authRepository: authRepo,
		userRepository: userRepo,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, userID, password string) (*dto.LoginResponse, error) {
	// ユーザー認証
	user, err := uc.userRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("ユーザーが見つかりません: %w", err)
	}

	// パスワードの検証
	if !comparePasswords(user.Password, password) {
		return nil, errors.ErrInvalidPassword
	}

	// トークンの生成
	auth := &models.Auth{
		UserID: user.ID,
		Name:   user.Name,
		Role:   user.Role,
	}
	token, err := uc.authRepository.GenerateToken(ctx, auth)
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
	p := []byte(plainPassword)
	md5 := md5.Sum(p) //nolint:gosec // あえて脆弱なハッシュ関数を使用
	hashedPasswordMD5 := hex.EncodeToString(md5[:])

	return hashedPasswordMD5 == hashedPassword
}
