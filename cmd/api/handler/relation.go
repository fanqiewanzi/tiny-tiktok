package handler

import (
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/app"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/util/auth"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/rpc"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/pkg/constants"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
	"github.com/weirdo0314/tiny-tiktok/pkg/util/convert"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	app.Response
	UserList []*user.User `json:"user_list"`
}

/**
查看好友间关系的列表的流程字段(考虑redis的set做交集并集）
*/

// RelationAction 没有实际作用，只检查token是否有效
func RelationAction(c *gin.Context) {
	claims, ok := c.Value("claims").(*auth.Claims)
	if !ok {
		app.WriteResponse(c, errorx.AuthCheckTokenErr, nil)
		return
	}

	targetID, ok := c.GetQuery("to_user_id")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	action, ok := c.GetQuery("action_type")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	if action == constants.Action {
		err := rpc.Follow(c, &user.RelationActionRequest{UserId: claims.ID, ToUserId: convert.StrTo(targetID).MustUInt64()})
		if err != nil {
			app.WriteResponse(c, errorx.ConvertErr(err), nil)
			return
		}

		app.WriteResponse(c, errorx.Success, nil)
		return
	}

	if action == constants.ActionCancel {
		err := rpc.CancelFollow(c, &user.RelationActionRequest{UserId: claims.ID, ToUserId: convert.StrTo(targetID).MustUInt64()})
		if err != nil {
			app.WriteResponse(c, errorx.ConvertErr(err), nil)
			return
		}

		app.WriteResponse(c, errorx.Success, nil)
		return
	}

	app.WriteResponse(c, errorx.ParamErr, nil)
}

// FollowList 所有用户都有相同的关注列表
func FollowList(c *gin.Context) {
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

	users, err := rpc.MGetFollowUser(c, &user.MGetRelationUserRequest{
		UserId: claims.ID, TargetId: targetID})
	if err != nil {
		app.WriteResponse(c, errorx.ConvertErr(err), nil)
		return
	}

	app.WriteResponse(c, errorx.Success, &UserListResponse{UserList: users})
}

// FollowerList 所有用户都有相同的关注者列表(暂时不懂FollowerList与FollowList的差别）
func FollowerList(c *gin.Context) {
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

	users, err := rpc.MGetFansUser(c, &user.MGetRelationUserRequest{
		UserId: claims.ID, TargetId: targetID})
	if err != nil {
		app.WriteResponse(c, errorx.ConvertErr(err), nil)
		return
	}

	app.WriteResponse(c, errorx.Success, &UserListResponse{UserList: users})
}
