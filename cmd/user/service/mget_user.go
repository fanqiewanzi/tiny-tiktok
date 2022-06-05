package service

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/user/dao"
	"github.com/weirdo0314/tiny-tiktok/cmd/user/pack"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
)

type MGetUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{
		ctx: ctx,
	}
}

func (s *MGetUserService) MGetUser(req *user.MGetUserRequest) ([]*user.User, error) {
	users, err := dao.MGetUsers(s.ctx, req.TargetUserIds)
	if err != nil {
		return nil, err
	}

	for i, v := range users {
		dao.QueryFollowInfo(s.ctx, users[i], req.UserId, uint64(v.ID))
	}

	return pack.Users(users), nil
}
