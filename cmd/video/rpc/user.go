package rpc

import (
	"context"
	"strings"
	"time"

	"github.com/weirdo0314/tiny-tiktok/cmd/video/config"
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

// MGetUser multiple get list of user info
func MGetUser(ctx context.Context, req *user.MGetUserRequest) (map[uint64]*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.BaseResp.StatusCode != 0 {
		return nil, errorx.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	res := make(map[uint64]*user.User)
	for _, u := range resp.Users {
		res[u.UserId] = u
	}

	return res, nil
}

// MGetUser multiple get list of user info
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
