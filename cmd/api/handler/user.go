package handler

import (
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/app"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/util/auth"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/rpc"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
	"github.com/weirdo0314/tiny-tiktok/pkg/util/convert"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	app.Response
	User *user.User `json:"user"`
}

// UserLoginResponse 用于登陆/注册返回的结构体
type UserLoginResponse struct {
	app.Response
	UserId uint64 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// Register 注册业务调用
func Register(c *gin.Context) {
	name, ok := c.GetQuery("username")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	password, ok := c.GetQuery("password")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	userID, err := rpc.CreateUser(c, &user.CreateUserRequest{UserName: name, Password: password})
	if err != nil {
		app.WriteResponse(c, errorx.ConvertErr(err), nil)
		return
	}

	token, err := auth.CreateToken(userID)
	if err != nil {
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, errorx.Success, &UserLoginResponse{UserId: userID, Token: token})
}

// Login 登录调用流程
func Login(c *gin.Context) {
	name, ok := c.GetQuery("username")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	password, ok := c.GetQuery("password")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	userID, err := rpc.CheckUser(c, &user.CheckUserRequest{UserName: name, Password: password})
	if err != nil {
		app.WriteResponse(c, err, nil)
		return
	}

	token, err := auth.CreateToken(userID)
	if err != nil {
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, errorx.Success, &UserLoginResponse{UserId: userID, Token: token})
}

func UserInfo(c *gin.Context) {
	claims, ok := c.Value("claims").(*auth.Claims)
	if !ok {
		app.WriteResponse(c, errorx.AuthCheckTokenErr, nil)
		return
	}

	userID, ok := c.GetQuery("user_id")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	var targetID uint64
	if userID == "0" {
		targetID = claims.ID
	} else {
		targetID = convert.StrTo(userID).MustUInt64()
	}

	user, err := rpc.GetUser(c, &user.GetUserRequest{UserId: claims.ID, TargetUserId: targetID})
	if err != nil {
		app.WriteResponse(c, errorx.ConvertErr(err), nil)
		return
	}

	app.WriteResponse(c, errorx.Success, &UserResponse{User: user})
}
