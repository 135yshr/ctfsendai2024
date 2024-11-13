package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenExpirationTime = time.Hour // トークンの有効期限
)

type JWTAuthRepository struct {
	secretKey string
	users     map[string]*models.Auth
}

type userConfig struct {
	Users []struct {
		ID       string `json:"id"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Role     string `json:"role"`
	} `json:"users"`
}

func NewJWTAuthRepository(secretKey string, configPath string) (repositories.AuthRepository, error) {
	repo := &JWTAuthRepository{
		secretKey: secretKey,
		users:     make(map[string]*models.Auth),
	}

	if err := repo.loadUsers(configPath); err != nil {
		return nil, fmt.Errorf("ユーザー情報の読み込みに失敗しました: %w", err)
	}

	return repo, nil
}

func (r *JWTAuthRepository) loadUsers(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("設定ファイルの読み込みに失敗しました: %w", err)
	}

	var config userConfig
	if err = json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("設定ファイルのパースに失敗しました: %w", err)
	}

	for _, user := range config.Users {
		r.users[user.ID] = &models.Auth{
			UserID:   user.ID,
			Name:     user.Name,
			Password: user.Password,
			Role:     models.Role(user.Role),
		}
	}

	return nil
}

// GenerateToken はJWTトークンを生成します.
func (r *JWTAuthRepository) GenerateToken(_ context.Context, auth *models.Auth) (*models.Token, error) {
	expiresAt := time.Now().Add(tokenExpirationTime)
	claims := jwt.MapClaims{
		"user_id": auth.UserID,
		"name":    auth.Name,
		"role":    auth.Role,
		"exp":     expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(r.secretKey))
	if err != nil {
		return nil, fmt.Errorf("トークンの生成に失敗しました: %w", err)
	}

	return &models.Token{
		AccessToken: signedToken,
		ExpiresAt:   expiresAt,
	}, nil
}

// ValidateToken はJWTトークンを検証します.
func (r *JWTAuthRepository) ValidateToken(_ context.Context, tokenString string) (*models.Auth, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidSignMethod
		}

		return []byte(r.secretKey), nil
	})
	if err != nil {
		return nil, errors.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.ErrInvalidToken
	}

	expirationTime, err := claims.GetExpirationTime()
	if err != nil || expirationTime.Before(time.Now()) {
		return nil, errors.ErrTokenExpired
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.ErrInvalidUserID
	}
	name, ok := claims["name"].(string)
	if !ok {
		return nil, errors.ErrInvalidName
	}
	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.ErrInvalidRole
	}

	return &models.Auth{UserID: userID, Name: name, Role: models.Role(role)}, nil
}

// FindByUsername はユーザー名からユーザーを検索します.
func (r *JWTAuthRepository) FindByUserID(_ context.Context, userID string) (*models.Auth, error) {
	if auth, exists := r.users[userID]; exists {
		return auth, nil
	}

	return nil, errors.ErrUserNotFound
}

// Store は認証情報を保存します.
func (r *JWTAuthRepository) Store(_ context.Context, auth *models.Auth) error {
	if auth == nil {
		return errors.ErrNilAuth
	}
	r.users[auth.UserID] = auth

	return nil
}
