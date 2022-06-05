package handler

import (
	"io/ioutil"
	"path/filepath"

	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/app"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/util/auth"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/rpc"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
	"github.com/weirdo0314/tiny-tiktok/pkg/util/convert"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	app.Response
	VideoList []*video.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	claims, ok := c.Value("claims").(*auth.Claims)
	if !ok {
		app.WriteResponse(c, errorx.AuthCheckTokenErr, nil)
		return
	}

	title, ok := c.GetPostForm("title")
	if !ok {
		app.WriteResponse(c, errorx.ParamErr, nil)
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		app.WriteResponse(c, err, nil)
		return
	}

	content, err := data.Open()
	if err != nil {
		app.WriteResponse(c, err, nil)
		return
	}

	byteFile, err := ioutil.ReadAll(content)
	if err != nil {
		app.WriteResponse(c, err, nil)
		return
	}

	err = rpc.Publish(c, &video.PublishActionRequest{
		UserId: claims.ID, Data: byteFile, Title: title, FileExt: filepath.Ext(data.Filename)})
	if err != nil {
		app.WriteResponse(c, errorx.ConvertErr(err), nil)
		return
	}

	app.WriteResponse(c, errorx.Success, nil)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
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

	videos, err := rpc.MGetVideo(c, &video.PublishListRequest{UserId: claims.ID, TargetId: targetID})
	if err != nil {
		app.WriteResponse(c, errorx.ConvertErr(err), nil)
		return
	}

	app.WriteResponse(c, errorx.Success, &VideoListResponse{VideoList: videos})
}
