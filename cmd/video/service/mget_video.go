package service

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/video/dao"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/pack"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/rpc"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
)

type MGetVideoService struct {
	ctx context.Context
}

func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

func (s *MGetVideoService) MGetVideos(req *video.PublishListRequest) ([]*video.Video, error) {
	videos, err := dao.MGetVideosByUID(s.ctx, req.TargetId)
	if err != nil {
		return nil, err
	}

	for i, v := range videos {
		videos[i].CommentCount, err = dao.QueryCommentCount(s.ctx, int64(v.ID))
		if err != nil {
			return nil, err
		}
		videos[i].FavoriteCount = dao.QueryFavoriteCount(s.ctx, uint64(v.ID))
		videos[i].IsFavorite = dao.QueryIsFavorite(s.ctx, req.UserId, uint64(v.ID))
	}

	user, err := rpc.GetUser(s.ctx, &user.GetUserRequest{UserId: req.UserId, TargetUserId: req.TargetId})
	if err != nil {
		return nil, err
	}

	return pack.VideosOfSameUser(videos, user), nil
}
