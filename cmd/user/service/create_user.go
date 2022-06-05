package service

import (
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/user/dao"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{
		ctx: ctx,
	}
}

func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) (uint64, error) {
	// 判断用户是否存在
	ok, err := NewCheckUserService(s.ctx).IsUserExist(req.UserName)
	// 操作数据库的过程会不会有err
	if err != nil {
		return 0, err
	}
	// 存在相同的用户返回
	if ok {
		return 0, errorx.UserAlreadyExistErr
	}

	// 通过bcrypt加密密码
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	req.Password = string(hash)

	// 数据库数据创建失败的情况
	id, err := dao.CreateUser(s.ctx, dao.User{Name: req.UserName, Password: req.Password})
	if err != nil {
		return 0, err
	}

	return id, nil
}
