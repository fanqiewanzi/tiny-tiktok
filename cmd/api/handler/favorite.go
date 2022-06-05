package handler

import (
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/app"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/util/auth"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/rpc"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
	"github.com/weirdo0314/tiny-tiktok/pkg/constants"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
	"github.com/weirdo0314/tiny-tiktok/pkg/util/convert"

	"github.com/gin-gonic/gin"
)

// FavoriteAction 没有实际作用，只检查token是否有效
func FavoriteAction(c *gin.Context) {
	claims, ok := c.Value("claims").(*auth.Claims)
	if !ok {
		app.WriteResponse(c, errorx.AuthCheckTokenErr, nil)
		return
	}

	vid, ok := c.GetQuery("video_id")
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
		err := rpc.Favorite(c, &video.FavoriteActionRequest{UserId: claims.ID, VideoId: convert.StrTo(vid).MustUInt64()})
		if err != nil {
			app.WriteResponse(c, errorx.ConvertErr(err), nil)
			return
		}

		app.WriteResponse(c, errorx.Success, nil)
		return
	}

	if action == constants.ActionCancel {
		err := rpc.CancelFavorite(c, &video.FavoriteActionRequest{UserId: claims.ID, VideoId: convert.StrTo(vid).MustUInt64()})
		if err != nil {
			app.WriteResponse(c, errorx.ConvertErr(err), nil)
		}

		app.WriteResponse(c, errorx.Success, nil)
		return
	}

	app.WriteResponse(c, errorx.ParamErr, nil)
}

// FavoriteList 所有用户都有相同的收藏视频列表
func FavoriteList(c *gin.Context) {
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

	videos, err := rpc.MGetFavorite(c, &video.FavoriteListRequest{UserId: claims.ID, TargetId: targetID})
	if err != nil {
		app.WriteResponse(c, errorx.ConvertErr(err), nil)
		return
	}

	app.WriteResponse(c, errorx.Success, &FeedResponse{VideoList: videos})
}
