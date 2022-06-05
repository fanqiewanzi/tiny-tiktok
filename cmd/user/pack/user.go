package pack

import (
	"github.com/weirdo0314/tiny-tiktok/cmd/user/dao"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
)

// User pack user info
func User(u *dao.User) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{UserId: uint64(u.ID), UserName: u.Name, FollowCount: u.FollowCount,
		FollowerCount: u.FollowerCount, IsFollow: u.IsFollow}
}

// Users pack list of user info
func Users(us []*dao.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
