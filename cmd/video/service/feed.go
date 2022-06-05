package service

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/video/dao"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/pack"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/rpc"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) Feed(req *video.FeedRequest) ([]*video.Video, error) {
	videos, err := dao.MGetVideos(s.ctx, req.UserId, uint64(req.LatestTime))
	if err != nil {
		return nil, err
	}

	userIDs := make([]uint64, len(videos))
	for i, v := range videos {
		videos[i].CommentCount, err = dao.QueryCommentCount(s.ctx, int64(v.ID))
		if err != nil {
			return nil, err
		}
		videos[i].FavoriteCount = dao.QueryFavoriteCount(s.ctx, uint64(v.ID))
		videos[i].IsFavorite = dao.QueryIsFavorite(s.ctx, req.UserId, uint64(v.ID))
		userIDs[i] = uint64(v.AuthorID)
	}

	users, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserId: req.UserId, TargetUserIds: userIDs})
	if err != nil {
		return nil, err
	}

	return pack.Videos(videos, users), nil
}
