package rpc

import (
	"context"
	"strings"
	"time"

	"github.com/weirdo0314/tiny-tiktok/cmd/api/config"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user/userservice"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
	"github.com/weirdo0314/tiny-tiktok/pkg/middleware"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver(strings.Split(config.Service.EtcdAddress, ";"))
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		config.Service.UserServiceName,
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
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *user.CreateUserRequest) (uint64, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.UserId, nil
}

// CheckUser check user info
func CheckUser(ctx context.Context, req *user.CheckUserRequest) (uint64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}

func MGetUser(ctx context.Context, req *user.MGetUserRequest) ([]*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.Users, nil
}

func GetUser(ctx context.Context, req *user.GetUserRequest) (*user.User, error) {
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.User, nil
}

func Follow(ctx context.Context, req *user.RelationActionRequest) error {
	resp, err := userClient.Follow(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errorx.NewError(resp.StatusCode, resp.StatusMessage)
	}

	return nil
}

func CancelFollow(ctx context.Context, req *user.RelationActionRequest) error {
	resp, err := userClient.CacelFollow(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errorx.NewError(resp.StatusCode, resp.StatusMessage)
	}

	return nil
}

func MGetFollowUser(ctx context.Context, req *user.MGetRelationUserRequest) ([]*user.User, error) {
	resp, err := userClient.MGetFollowUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.Users, nil
}

func MGetFansUser(ctx context.Context, req *user.MGetRelationUserRequest) ([]*user.User, error) {
	resp, err := userClient.MGetFansUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.Users, nil
}
