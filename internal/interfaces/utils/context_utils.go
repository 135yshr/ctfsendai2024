package utils

import (
	"github.com/135yshr/ctfsendai2024/internal/domain/models"
	"github.com/gin-gonic/gin"
)

const adminUserKey = "user"

// GetUserFromContext はコンテキストからユーザー情報を取得します.
func GetUserFromContext(c *gin.Context) (*models.Auth, bool) {
	user, exists := c.Get(adminUserKey)
	if !exists {
		return nil, false
	}
	auth, ok := user.(*models.Auth)
	if !ok {
		return nil, false
	}

	return auth, true
}

func SetAdminUserToContext(c *gin.Context, auth *models.Auth) {
	c.Set(adminUserKey, auth)
}
