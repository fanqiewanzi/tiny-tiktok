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

type CommentListResponse struct {
	app.Response
	CommentList []*video.Comment `json:"comment_list,omitempty"`
}

type CommentResponse struct {
	app.Response
	Comment *video.Comment
}

// CommentAction comment video and return a comment struct if success
func CommentAction(c *gin.Context) {
	claims, ok := c.Value("claims").(*auth.Claims)
	if !ok {
		app.WriteResponse(c, errorx.AuthCheckTokenErr, nil)
		return
	}

	action, ok := c.GetQuery("action_type")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	vid, ok := c.GetQuery("video_id")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	if action == constants.Action {
		content, ok := c.GetQuery("comment_text")
		if !ok {
			app.WriteResponse(c, errorx.ParamErr, nil)
			return
		}

		comment, err := rpc.Comment(c, &video.CommentActionRequest{
			UserId: claims.ID, VideoId: convert.StrTo(vid).MustUInt64(), CommentText: content})
		if err != nil {
			app.WriteResponse(c, errorx.ConvertErr(err), &CommentResponse{Comment: comment})
			return
		}

		app.WriteResponse(c, errorx.Success, &CommentResponse{Comment: comment})
		return
	}

	if action == constants.ActionCancel {
		cid, ok := c.GetQuery("comment_id")
		if !ok {
			app.WriteResponse(c, errorx.ParamErr, nil)
			return
		}

		err := rpc.DeleteComment(c, &video.CommentActionRequest{UserId: claims.ID, CommentId: convert.StrTo(cid).MustUInt64()})
		if err != nil {
			app.WriteResponse(c, errorx.ConvertErr(err), nil)
			return
		}

		app.WriteResponse(c, errorx.Success, nil)
		return
	}

	app.WriteResponse(c, errorx.ParamErr, nil)
}

// CommentList get all comments of one video
func CommentList(c *gin.Context) {
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

	vidoes, err := rpc.MGetComment(c, &video.CommentListRequest{UserId: claims.ID, VideoId: convert.StrTo(vid).MustUInt64()})
	if err != nil {
		app.WriteResponse(c, errorx.ConvertErr(err), nil)
		return
	}

	app.WriteResponse(c, errorx.Success, &CommentListResponse{CommentList: vidoes})
}
