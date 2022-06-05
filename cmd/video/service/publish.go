package service

import (
	"context"
	"path/filepath"

	"github.com/weirdo0314/tiny-tiktok/cmd/video/config"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/dao"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/util/oss"
	"github.com/weirdo0314/tiny-tiktok/kitex_gen/video"
	"github.com/weirdo0314/tiny-tiktok/pkg/util/filename"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{
		ctx: ctx,
	}
}

func (s *PublishService) Publish(req *video.PublishActionRequest) error {
	name := filename.RandFileName()
	err := oss.Upload(s.ctx, req.Data, "video/"+name+req.FileExt)
	if err != nil {
		return err
	}

	return dao.CreateVideo(s.ctx, []*dao.Video{{
		AuthorID: req.UserId,
		Title:    req.Title,
		PlayUrl:  config.Service.URL + filepath.Join("/video/", name+req.FileExt),
		CoverUrl: config.Service.URL + filepath.Join("/image/", name+".jpg")}})
}
