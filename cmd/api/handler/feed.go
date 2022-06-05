package handler

import (
	"time"

	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/app"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/util/auth"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/rpc"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
	"github.com/weirdo0314/tiny-tiktok/pkg/util/convert"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	app.Response
	VideoList []*video.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var id uint64
	claims, ok := c.Value("claims").(*auth.Claims)
	if ok {
		id = claims.ID
	}

	lastTime, ok := c.GetQuery("latest_time")
	var nextTime int64
	if ok {
		nextTime = convert.StrTo(lastTime[:10]).MustInt64()
	} else {
		nextTime = time.Now().Unix()
	}

	videos, err := rpc.Feed(c, &video.FeedRequest{UserId: id, LatestTime: nextTime})
	if err != nil {
		app.WriteResponse(c, errorx.ConvertErr(err), nil)
		return
	}

	app.WriteResponse(c, errorx.Success, &FeedResponse{VideoList: videos, NextTime: time.Now().Unix()})
}
