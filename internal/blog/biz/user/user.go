package user

import (
	"blog/internal/blog/store"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/model"
	"blog/pkg/api/blog"
	"context"
	"errors"
)

type UserBiz interface {
	Create(ctx context.Context, r *blog.CreateUserRequest) error
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IStore) UserBiz {
	return &userBiz{ds: ds}
}

func (b *userBiz) Create(ctx context.Context, r *blog.CreateUserRequest) error {
	userm := model.UserM{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
		Nickname: r.Nickname,
		Phone:    r.Phone,
	}

	if err := b.ds.Users().Create(ctx, &userm); err != nil {
		if errors.Is(err, errno.ErrUserAlreadyExist) {
			return errno.ErrUserAlreadyExist
		}
		return err
	}

	return nil
}
