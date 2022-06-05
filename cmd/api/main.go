package main

import (
	"log"
	"net/http"

	"github.com/weirdo0314/tiny-tiktok/cmd/api/config"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/handler"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/rpc"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config.Init()
	rpc.Init()
	log.Print(config.Service)
	defer zap.L().Sync()
	gin.SetMode(config.Service.RunMode)
	r := gin.Default()
	handler.InitRouter(r)
	s := &http.Server{
		Addr:           ":" + config.Service.HttpPort,
		Handler:        r,
		ReadTimeout:    config.Service.ReadTimeout,  //允许读取的最大时间
		WriteTimeout:   config.Service.WriteTimeout, //允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                     //允许请求头的最大字节数
	}
	log.Fatal(s.ListenAndServe())
}
