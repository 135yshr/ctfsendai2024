package middleware

import (
	"net/http"
	"strings"

	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authRepo repositories.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
			c.Abort()

			return
		}

		// トークン検証ロジック
		auth, err := authRepo.ValidateToken(c.Request.Context(), strings.Replace(token, "Bearer ", "", 1))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "無効なトークンです"})
			c.Abort()

			return
		}

		c.Set("user", auth)
		c.Next()
	}
}
