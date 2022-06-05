package service

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/user/dao"
	"github.com/weirdo0314/tiny-tiktok/cmd/user/pack"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{
		ctx: ctx,
	}
}

func (s *GetUserService) GetUser(req *user.GetUserRequest) (*user.User, error) {
	user, err := dao.QueryUser(s.ctx, req.UserId, req.TargetUserId)
	if err != nil {
		return nil, err
	}

	dao.QueryFollowInfo(s.ctx, user, req.UserId, req.TargetUserId)
	return pack.User(user), nil
}
