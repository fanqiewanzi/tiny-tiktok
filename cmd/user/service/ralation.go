package service

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/user/dao"
	"github.com/weirdo0314/tiny-tiktok/cmd/user/pack"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
)

type RelationService struct {
	ctx context.Context
}

// NewFollowService new FollowService
func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}

func (s *RelationService) Follow(req *user.RelationActionRequest) error {
	return dao.CreateFollow(s.ctx, req.UserId, req.ToUserId)
}

func (s *RelationService) CancelFollow(req *user.RelationActionRequest) error {
	return dao.DeleteFollow(s.ctx, req.UserId, req.ToUserId)
}

func (s *RelationService) GetFollowOrFans(req *user.MGetRelationUserRequest, key string) ([]*user.User, error) {
	ids, err := dao.MGetFollowOrFansIDs(s.ctx, key, req.TargetId)
	if err != nil {
		return nil, err
	}

	users, err := dao.MGetUsers(s.ctx, ids)
	if err != nil {
		return nil, err
	}

	for i, v := range users {
		dao.QueryFollowInfo(s.ctx, users[i], req.UserId, uint64(v.ID))
	}

	return pack.Users(users), nil
}
