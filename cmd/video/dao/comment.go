package dao

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/pkg/constants"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UID     uint64 `gorm:"column:uid"`
	VID     uint64 `gorm:"column:vid"`
	Content string `gorm:"column:content"`
}

func (m *Comment) TableName() string {
	return constants.CommentTableName
}

func QueryCommentCount(ctx context.Context, vid int64) (uint64, error) {
	var count int64
	if err := db.WithContext(ctx).Model(&Comment{}).Where("vid = ?", vid).Count(&count).Error; err != nil {
		return 0, err
	}

	return uint64(count), nil
}

func CreateComment(ctx context.Context, comment *Comment) error {
	return db.WithContext(ctx).Create(comment).Error
}

func DeleteComment(ctx context.Context, uid, cid uint64) error {
	return db.WithContext(ctx).Where("uid = ? and cid = ?", uid, cid).Delete(&Comment{}).Error
}

func MGetComments(ctx context.Context, vid uint64) ([]*Comment, error) {
	var comments []*Comment
	if err := db.WithContext(ctx).Where("vid = ?", vid).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}
