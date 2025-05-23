package model

import (
	"blog/pkg/util/id"
	"time"

	"gorm.io/gorm"
)

type PostM struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username;not null"`
	PostID    string    `gorm:"column:postID;not null"`
	Title     string    `gorm:"column:title;not null"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

func (p *PostM) TableName() string {
	return "post"
}

func (p *PostM) BeforeCreate(tx *gorm.DB) error {
	p.PostID = "post-" + id.GenShortID()

	return nil
}
