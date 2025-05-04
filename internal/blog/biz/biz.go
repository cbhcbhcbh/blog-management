package biz

import (
	"blog/internal/blog/biz/post"
	"blog/internal/blog/biz/user"
	"blog/internal/blog/store"
)

type IBiz interface {
	Users() user.UserBiz
	Posts() post.PostBiz
}

type biz struct {
	ds store.IStore
}

var _ IBiz = (*biz)(nil)

func NewBiz(ds store.IStore) IBiz {
	return &biz{ds: ds}
}

func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}

func (b *biz) Posts() post.PostBiz {
	return post.New(b.ds)
}
