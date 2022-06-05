package oss

import (
	"bytes"
	"context"

	"github.com/weirdo0314/tiny-tiktok/cmd/video/config"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func Upload(ctx context.Context, data []byte, fileName string) error {
	putPolicy := storage.PutPolicy{
		Scope: config.Service.Bucket,
	}
	mac := qbox.NewMac(config.Service.AccessKey, config.Service.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	//构建上传表单对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.Put(ctx, &ret, upToken, fileName, bytes.NewReader(data), int64(len(data)), nil)
	if err != nil {
		return err
	}

	return nil
}
