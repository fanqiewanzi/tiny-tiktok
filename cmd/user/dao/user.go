package dao

import (
	"context"
	"errors"

	"github.com/weirdo0314/tiny-tiktok/pkg/constants"

	"gorm.io/gorm"
)

// User 用户的字段
type User struct {
	gorm.Model
	Name          string `gorm:"column:name"`
	Password      string `gorm:"column:password"`
	FollowCount   uint64 `gorm:"-"`
	FollowerCount uint64 `gorm:"-"`
	IsFollow      bool   `gorm:"-"`
}

func (m *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, user User) (uint64, error) {
	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		return 0, err
	}
	return uint64(user.ID), nil
}

// GetUserByName get user info by username
func GetUserByName(ctx context.Context, username string) (*User, error) {
	var user User
	if err := db.WithContext(ctx).Where("name = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &user, nil
		}
		return &user, err
	}

	return &user, nil
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []uint64) ([]*User, error) {
	var res []*User
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := db.WithContext(ctx).Select("id", "name").Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

// QueryUser query user info
func QueryUser(ctx context.Context, curID, targetID uint64) (*User, error) {
	var user User
	if err := db.WithContext(ctx).Where("id = ?", targetID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}
		return nil, err
	}

	return &user, nil
}

func QueryFollowInfo(ctx context.Context, user *User, uid, targetID uint64) {
	user.FollowCount = QueryFollowCount(ctx, targetID)
	user.FollowerCount = QueryFollowerCount(ctx, targetID)
	user.IsFollow = QueryIsFollow(ctx, uid, targetID)
}
