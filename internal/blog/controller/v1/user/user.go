package user

import (
	"blog/internal/blog/biz"
	"blog/internal/blog/store"
	"blog/pkg/auth"

	pb "blog/pkg/proto/blog/v1"
)

type UserController struct {
	a *auth.Authz
	b biz.IBiz
	pb.UnimplementedBlogServer
}

func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{
		a: a,
		b: biz.NewBiz(ds),
	}
}
