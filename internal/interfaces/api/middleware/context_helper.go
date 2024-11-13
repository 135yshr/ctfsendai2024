package middleware

import (
	"context"

	"github.com/135yshr/ctfsendai2024/internal/domain/errors"
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
)

type userContextKey struct{}

// GetUserFromContext はコンテキストからユーザー情報を取得します.
func GetUserFromContext(ctx context.Context) (*models.Auth, error) {
	value, ok := ctx.Value(userContextKey{}).(*models.Auth)
	if !ok {
		return nil, errors.ErrInvalidUser
	}

	return value, nil
}

func SetUserToContext(ctx context.Context, user *models.Auth) context.Context {
	return context.WithValue(ctx, userContextKey{}, user)
}
