package store

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"blog/internal/pkg/model"
)

type PostStore interface {
	Create(ctx context.Context, post *model.PostM) error
	Get(ctx context.Context, username, postID string) (*model.PostM, error)
	Update(ctx context.Context, post *model.PostM) error
	List(ctx context.Context, username string, offset, limit int) (int64, []*model.PostM, error)
	Delete(ctx context.Context, username string, postIDs []string) error
}

type posts struct {
	db *gorm.DB
}

var _ PostStore = (*posts)(nil)

func newPosts(db *gorm.DB) *posts {
	return &posts{db}
}

func (u *posts) Create(ctx context.Context, post *model.PostM) error {
	return u.db.Create(&post).Error
}

func (u *posts) Get(ctx context.Context, username, postID string) (*model.PostM, error) {
	var post model.PostM
	if err := u.db.Where("username = ? and postID = ?", username, postID).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (u *posts) Update(ctx context.Context, post *model.PostM) error {
	return u.db.Save(post).Error
}

func (u *posts) List(ctx context.Context, username string, offset, limit int) (count int64, ret []*model.PostM, err error) {
	err = u.db.Where("username = ?", username).Offset(offset).Limit(defaultLimit(limit)).Order("id desc").Find(&ret).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error

	return
}

func (u *posts) Delete(ctx context.Context, username string, postIDs []string) error {
	err := u.db.Where("username = ? and postID in (?)", username, postIDs).Delete(&model.PostM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
