package dao

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/pkg/constants"

	"gorm.io/gorm"
)

// Video 影音视频字段
type Video struct {
	gorm.Model
	Title         string `gorm:"column:title"`
	AuthorID      uint64 `gorm:"column:uid"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount uint64 `gorm:"-"`
	CommentCount  uint64 `gorm:"-"`
	IsFavorite    bool   `gorm:"-"`
}

func (m *Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	if err := db.WithContext(ctx).Create(videos).Error; err != nil {
		return err
	}

	return nil
}

// MGetVideos multiple get list of video info
func MGetVideos(ctx context.Context, uid, nextTime uint64) ([]*Video, error) {
	var res []*Video
	if err := db.WithContext(ctx).Debug().Order("created_at DESC").Where("unix_timestamp(created_at) < ?", nextTime).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func MGetVideosByUID(ctx context.Context, uid uint64) ([]*Video, error) {
	var res []*Video
	if err := db.WithContext(ctx).Where("uid = ?", uid).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func MGetVideosByIDs(ctx context.Context, vids []uint64) ([]*Video, error) {
	var res []*Video
	if err := db.WithContext(ctx).Where("id in ?", vids).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
