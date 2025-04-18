package user

import (
	"blog/internal/blog/biz"
	"blog/internal/blog/store"
)

type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{
		b: biz.NewBiz(ds),
	}
}
