package service

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/video/dao"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/pack"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/rpc"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{
		ctx: ctx,
	}
}

func (s *CommentService) Comment(req *video.CommentActionRequest) (*video.Comment, error) {
	comment := &dao.Comment{UID: req.UserId, VID: req.VideoId, Content: req.CommentText}
	err := dao.CreateComment(s.ctx, comment)
	if err != nil {
		return nil, err
	}

	user, err := rpc.GetUser(s.ctx, &user.GetUserRequest{UserId: req.UserId, TargetUserId: req.UserId})
	if err != nil {
		return nil, err
	}

	return pack.Comment(comment, user), nil
}

func (s *CommentService) DeleteComment(req *video.CommentActionRequest) error {
	return dao.DeleteComment(s.ctx, req.UserId, req.CommentId)
}

func (s *CommentService) MGetComment(req *video.CommentListRequest) ([]*video.Comment, error) {
	comments, err := dao.MGetComments(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	userIDs := make([]uint64, len(comments))
	for i, v := range comments {
		userIDs[i] = v.UID
	}

	users, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserId: req.UserId, TargetUserIds: userIDs})
	if err != nil {
		return nil, err
	}

	return pack.Comments(comments, users), nil
}
