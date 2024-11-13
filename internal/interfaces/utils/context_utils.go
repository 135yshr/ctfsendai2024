package utils

import (
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// GetUserFromContext はコンテキストからユーザー情報を取得します.
func GetUserFromContext(c *gin.Context) (*models.Auth, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	auth, ok := user.(*models.Auth)
	if !ok {
		return nil, false
	}

	return auth, true
}
