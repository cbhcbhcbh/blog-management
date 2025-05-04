package post

import (
	"blog/internal/blog/biz"
	"blog/internal/blog/store"
)

type PostController struct {
	b biz.IBiz
}

func New(ds store.IStore) *PostController {
	return &PostController{b: biz.NewBiz(ds)}
}
