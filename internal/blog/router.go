package blog

import (
	"blog/internal/blog/controller/v1/post"
	"blog/internal/blog/controller/v1/user"
	"blog/internal/blog/store"
	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/log"
	mw "blog/internal/pkg/middleware"
	"blog/pkg/auth"

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

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)
	pc := post.New(store.S)

	router.POST("/login", uc.Login)

	v1 := router.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mw.Authn(), mw.Authz(authz))
			userv1.GET(":name", uc.Get)
			userv1.PUT(":name", uc.Update)
			userv1.GET("", uc.List)
			userv1.DELETE(":name", uc.Delete)
		}

		postv1 := v1.Group("/posts", mw.Authn())
		{
			postv1.POST("", pc.Create)
			postv1.GET(":postID", pc.Get)
			postv1.PUT(":postID", pc.Update)
			postv1.DELETE("", pc.DeleteCollection)
			postv1.GET("", pc.List)
			postv1.DELETE(":postID", pc.Delete)
		}
	}

	return nil
}
