package user

import (
	"blog/internal/blog/store"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/model"
	v1 "blog/pkg/api/blog/v1"
	"blog/pkg/auth"
	"blog/pkg/token"
	"context"
	"errors"
)

type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IStore) UserBiz {
	return &userBiz{ds: ds}
}

func (b *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
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

func (b *userBiz) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	user, err := b.ds.Users().Get(ctx, r.Username)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}

	if err := auth.Compare(user.Password, r.Password); err != nil {
		return nil, errno.ErrPasswordIncorrect
	}

	t, err := token.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, errno.ErrSignToken
	}

	return &v1.LoginResponse{Token: t}, nil
}
