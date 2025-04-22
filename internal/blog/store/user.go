package store

import (
	"blog/internal/pkg/errno"
	"blog/internal/pkg/model"
	"context"

	"gorm.io/gorm"
)

type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
	Get(ctx context.Context, username string) (*model.UserM, error)
}

type users struct {
	db *gorm.DB
}

var _ UserStore = (*users)(nil)

func NewUser(db *gorm.DB) *users {
	return &users{db: db}
}

func (u *users) Create(ctx context.Context, user *model.UserM) error {
	var count int64
	if err := u.db.WithContext(ctx).Model(&model.UserM{}).Where("username = ?", user.Username).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errno.ErrUserAlreadyExist
	}

	return u.db.WithContext(ctx).Create(user).Error
}

func (u *users) Get(ctx context.Context, username string) (*model.UserM, error) {
	var user model.UserM
	if err := u.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
