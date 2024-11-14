package middleware

import (
	"bytes"
	"io"

	"github.com/135yshr/ctfsendai2024/internal/foundation/logger"
	"github.com/gin-gonic/gin"
)

func RequestLogger(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		logger.Info("incoming request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"body", string(bodyBytes),
		)

		c.Next()
	}
}
