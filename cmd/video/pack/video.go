package pack

import (
	"github.com/weirdo0314/tiny-tiktok/cmd/video/dao"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
)

// User pack video info
func Video(v *dao.Video, u *user.User) *video.Video {
	if v == nil {
		return nil
	}

	return &video.Video{Id: uint64(v.ID), Title: v.Title, PlayUrl: v.PlayUrl, CoverUrl: v.CoverUrl,
		Author: u, FavoriteCount: v.FavoriteCount, CommentCount: v.CommentCount, IsFavorite: v.IsFavorite}
}

// Videos pack list of video info
func Videos(vs []*dao.Video, users map[uint64]*user.User) []*video.Video {
	videos := make([]*video.Video, 0, len(users))
	for _, v := range vs {
		if video := Video(v, users[uint64(v.AuthorID)]); video != nil {
			videos = append(videos, video)
		}
	}
	return videos
}

func VideosOfSameUser(vs []*dao.Video, user *user.User) []*video.Video {
	videos := make([]*video.Video, 0)
	for _, v := range vs {
		if video2 := Video(v, user); video2 != nil {
			videos = append(videos, video2)
		}
	}
	return videos
}
