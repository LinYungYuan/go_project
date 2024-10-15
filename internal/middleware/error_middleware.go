// internal/middleware/error_middleware.go
package middleware

import (
	"go_project/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				logger.Error("Request error",
					zap.String("error", e.Error()),
					zap.String("path", c.Request.URL.Path),
				)
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": c.Errors.Errors(),
			})
		}
	}
}
