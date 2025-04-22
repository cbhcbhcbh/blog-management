package blog

import (
	"blog/internal/blog/controller/v1/user"
	"blog/internal/blog/store"
	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

func installRouters(router *gin.Engine) error {
	router.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	router.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S)

	router.POST("/login", uc.Login)

	v1 := router.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
		}
	}

	return nil
}
