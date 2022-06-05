package pack

import (
	"errors"
	"time"

	"github.com/weirdo0314/tiny-tiktok/kitex_gen/user"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *user.BaseResponse {
	if err == nil {
		return baseResp(errorx.Success)
	}

	e := errorx.Error{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errorx.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errorx.Error) *user.BaseResponse {
	return &user.BaseResponse{StatusCode: err.Code, StatusMessage: err.Msg, ServiceTime: time.Now().Unix()}
}
