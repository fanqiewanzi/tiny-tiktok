package service

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/user/dao"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"

	"golang.org/x/crypto/bcrypt"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (uint64, error) {
	user, err := dao.GetUserByName(s.ctx, req.UserName)
	if err != nil {
		return 0, errorx.LoginErr
	}

	if user.ID == 0 {
		return 0, errorx.LoginErr
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return 0, errorx.LoginErr
	}

	return uint64(user.ID), nil
}

func (s *CheckUserService) IsUserExist(username string) (bool, error) {
	user, err := dao.GetUserByName(s.ctx, username)
	if err != nil {
		return false, err
	}

	return user.ID > 0, nil
}
