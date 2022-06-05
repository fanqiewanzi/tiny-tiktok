package pack

import (
	"github.com/weirdo0314/tiny-tiktok/cmd/video/dao"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
)

// Comment pack comment info
func Comment(c *dao.Comment, u *user.User) *video.Comment {
	if c == nil {
		return nil
	}

	return &video.Comment{Id: uint64(c.ID), User: u, Content: c.Content, CreateDate: c.CreatedAt.String()}
}

// Comments pack list of comment info
func Comments(cs []*dao.Comment, users map[uint64]*user.User) []*video.Comment {
	comments := make([]*video.Comment, 0, len(users))
	for _, v := range cs {
		if comment := Comment(v, users[uint64(v.UID)]); comment != nil {
			comments = append(comments, comment)
		}
	}
	return comments
}
