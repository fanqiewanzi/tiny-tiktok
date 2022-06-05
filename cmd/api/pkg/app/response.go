package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
)

type Responser interface {
	SetCode(int)
	SetMsg(string)
}

// Response 响应结构体
type Response struct {
	Code int    `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func (r *Response) SetCode(code int) {
	r.Code = code
}

func (r *Response) SetMsg(msg string) {
	r.Msg = msg
}

// WriteResponse 写入返回结果
func WriteResponse(c *gin.Context, err error, data Responser) {
	if data == nil {
		data = new(Response)
	}
	log.Println(data)
	Err := errorx.ConvertErr(err)
	data.SetMsg(Err.Msg)
	data.SetCode(int(Err.Code))
	c.JSON(http.StatusOK, data)
}
