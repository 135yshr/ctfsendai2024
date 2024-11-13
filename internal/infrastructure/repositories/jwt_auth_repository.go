package repositories

import (
	"fmt"
	"time"

	"github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthRepository struct {
	secretKey string
	users     map[string]*models.Auth
}

func NewJWTAuthRepository(secretKey string) repositories.AuthRepository {
	return &JWTAuthRepository{
		secretKey: secretKey,
		users:     make(map[string]*models.Auth),
	}
}

// GenerateToken はJWTトークンを生成します.
func (r *JWTAuthRepository) GenerateToken(auth *models.Auth) (*models.Token, error) {
	expiresAt := time.Now().Add(time.Hour)
	claims := jwt.MapClaims{
		"user_id": auth.UserID,
		"exp":     expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(r.secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &models.Token{
		AccessToken: signedToken,
		ExpiresAt:   expiresAt,
	}, nil
}

// ValidateToken はJWTトークンを検証します.
func (r *JWTAuthRepository) ValidateToken(tokenString string) (*models.Auth, error) {
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

	return &models.Auth{UserID: userID}, nil
}

// FindByUsername はユーザー名からユーザーを検索します.
func (r *JWTAuthRepository) FindByUsername(username string) (*models.Auth, error) {
	if auth, exists := r.users[username]; exists {
		return auth, nil
	}

	return nil, errors.ErrUserNotFound
}

// Store は認証情報を保存します.
func (r *JWTAuthRepository) Store(auth *models.Auth) error {
	if auth == nil {
		return errors.ErrNilAuth
	}
	r.users[auth.Username] = auth

	return nil
}
