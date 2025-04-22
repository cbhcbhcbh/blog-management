package middleware

import (
	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/known"
	"blog/pkg/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		claims, err := token.ParseToken(parts[1])
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set(known.XUsernameKey, claims.Username)

		c.Next()
	}
}
