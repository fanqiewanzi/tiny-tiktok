package main

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/user/pack"
	"github.com/weirdo0314/tiny-tiktok/cmd/user/service"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/pkg/constants"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)
	if len(req.UserName) == 0 || len(req.Password) == 0 || len(req.UserName) > 32 || len(req.Password) > 32 {
		resp.BaseResp = pack.BuildBaseResp(errorx.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	resp.UserId = uid
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	resp = new(user.MGetUserResponse)
	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	resp.Users = users
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	resp = new(user.GetUserResponse)
	user, err := service.NewGetUserService(ctx).GetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	resp.User = user
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	resp = new(user.CheckUserResponse)
	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errorx.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	return resp, nil
}

// Follow implements the UserServiceImpl interface.
func (s *UserServiceImpl) Follow(ctx context.Context, req *user.RelationActionRequest) (resp *user.BaseResponse, err error) {
	err = service.NewRelationService(ctx).Follow(req)
	if err != nil {
		resp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp = pack.BuildBaseResp(errorx.Success)
	return resp, nil
}

// CacelFollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) CacelFollow(ctx context.Context, req *user.RelationActionRequest) (resp *user.BaseResponse, err error) {
	err = service.NewRelationService(ctx).CancelFollow(req)
	if err != nil {
		resp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp = pack.BuildBaseResp(errorx.Success)
	return resp, nil
}

// MGetFollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetFollowUser(ctx context.Context, req *user.MGetRelationUserRequest) (resp *user.MGetRelationUserResponse, err error) {
	resp = new(user.MGetRelationUserResponse)
	users, err := service.NewRelationService(ctx).GetFollowOrFans(req, constants.FollowKey)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Users = users
	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	return resp, nil
}

// MGetFansUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetFansUser(ctx context.Context, req *user.MGetRelationUserRequest) (resp *user.MGetRelationUserResponse, err error) {
	resp = new(user.MGetRelationUserResponse)
	users, err := service.NewRelationService(ctx).GetFollowOrFans(req, constants.FansKey)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Users = users
	resp.BaseResp = pack.BuildBaseResp(errorx.Success)
	return resp, nil
}
