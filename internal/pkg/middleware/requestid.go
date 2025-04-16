package middleware

import (
	"blog/internal/pkg/known"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(known.XRequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set(known.XRequestIDKey, requestID)

		c.Writer.Header().Set(known.XRequestIDKey, requestID)
		c.Next()
	}
}
