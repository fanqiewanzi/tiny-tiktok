package rpc

import (
	"context"
	"strings"
	"time"

	"github.com/weirdo0314/tiny-tiktok/cmd/api/config"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video/videoservice"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
	"github.com/weirdo0314/tiny-tiktok/pkg/middleware"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver(strings.Split(config.Service.EtcdAddress, ";"))
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		config.Service.VideoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func Feed(ctx context.Context, req *video.FeedRequest) ([]*video.Video, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.VideoList, nil
}

func Publish(ctx context.Context, req *video.PublishActionRequest) error {
	resp, err := videoClient.Publish(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errorx.NewError(resp.StatusCode, resp.StatusMessage)
	}

	return nil
}

func MGetVideo(ctx context.Context, req *video.PublishListRequest) ([]*video.Video, error) {
	resp, err := videoClient.MGetVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.VideoList, nil
}

func Favorite(ctx context.Context, req *video.FavoriteActionRequest) error {
	resp, err := videoClient.Favorite(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errorx.NewError(resp.StatusCode, resp.StatusMessage)
	}

	return nil
}

func CancelFavorite(ctx context.Context, req *video.FavoriteActionRequest) error {
	resp, err := videoClient.CancelFavorite(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errorx.NewError(resp.StatusCode, resp.StatusMessage)
	}

	return nil
}

func MGetFavorite(ctx context.Context, req *video.FavoriteListRequest) ([]*video.Video, error) {
	resp, err := videoClient.MGetFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.VideoList, nil
}

func Comment(ctx context.Context, req *video.CommentActionRequest) (*video.Comment, error) {
	resp, err := videoClient.Comment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.Comment, nil
}

func DeleteComment(ctx context.Context, req *video.CommentActionRequest) error {
	resp, err := videoClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

func MGetComment(ctx context.Context, req *video.CommentListRequest) ([]*video.Comment, error) {
	resp, err := videoClient.MGetComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.CommentList, nil
}
