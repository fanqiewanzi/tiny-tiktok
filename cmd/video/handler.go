package main

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/video/pack"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/service"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(video.FeedResponse)
	videos, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	resp.VideoList = videos
	return resp, nil
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, req *video.PublishActionRequest) (resp *user.BaseResponse, err error) {
	err = service.NewPublishService(ctx).Publish(req)
	if err != nil {
		resp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp = pack.BuildBaseResp(errorx.Success)
	return resp, nil
}

// MGetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideo(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	resp = new(video.PublishListResponse)
	videos, err := service.NewMGetVideoService(ctx).MGetVideos(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	resp.VideoList = videos
	return resp, nil
}

// Favorite implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Favorite(ctx context.Context, req *video.FavoriteActionRequest) (resp *user.BaseResponse, err error) {
	err = service.NewFavoriteSevrice(ctx).Favorite(req)
	if err != nil {
		resp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp = pack.BuildBaseResp(errorx.Success)
	return resp, nil
}

// CancelFavorite implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CancelFavorite(ctx context.Context, req *video.FavoriteActionRequest) (resp *user.BaseResponse, err error) {
	err = service.NewFavoriteSevrice(ctx).CancelFavorite(req)
	if err != nil {
		resp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp = pack.BuildBaseResp(errorx.Success)
	return resp, nil
}

// MGetFavorite implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetFavorite(ctx context.Context, req *video.FavoriteListRequest) (resp *video.FavoriteListResponse, err error) {
	resp = new(video.FavoriteListResponse)
	videos, err := service.NewFavoriteSevrice(ctx).MGetFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	resp.VideoList = videos
	return resp, nil
}

// Comment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Comment(ctx context.Context, req *video.CommentActionRequest) (resp *video.CommentActionResponse, err error) {
	resp = new(video.CommentActionResponse)
	comment, err := service.NewCommentService(ctx).Comment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	resp.Comment = comment
	return resp, nil
}

// DeleteComment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DeleteComment(ctx context.Context, req *video.CommentActionRequest) (resp *video.CommentActionResponse, err error) {
	resp = new(video.CommentActionResponse)
	err = service.NewCommentService(ctx).DeleteComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	return resp, nil
}

// MGetComment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetComment(ctx context.Context, req *video.CommentListRequest) (resp *video.CommentListResponse, err error) {
	resp = new(video.CommentListResponse)
	comments, err := service.NewCommentService(ctx).MGetComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	resp.CommentList = comments
	return resp, nil
}
