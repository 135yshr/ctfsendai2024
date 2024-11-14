package middleware

import (
	"time"

	"github.com/135yshr/ctfsendai2024/internal/foundation/logger"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// レイテンシーを計算
		latency := time.Since(start)

		// 構造化ログを出力
		logger.Info("access_log",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"client_ip", c.ClientIP(),
			"status", c.Writer.Status(),
			"latency", latency,
			"timestamp", time.Now().Format(time.RFC3339),
		)
	}
}
