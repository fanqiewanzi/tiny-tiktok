package service

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/user/dao"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
)

type CancelFollowService struct {
	ctx context.Context
}

// NewCancelFollowService new CancelFollowService
func NewCancelFollowService(ctx context.Context) *CancelFollowService {
	return &CancelFollowService{
		ctx: ctx,
	}
}

func (s *CancelFollowService) CancelFollow(req *user.RelationActionRequest) error {
	return dao.DeleteFollow(s.ctx, req.UserId, req.ToUserId)
}
