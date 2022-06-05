package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/weirdo0314/tiny-tiktok/cmd/video/config"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/dao"
	"github.com/weirdo0314/tiny-tiktok/cmd/video/rpc"
	video "github.com/weirdo0314/tiny-tiktok/kitex_gen/video/videoservice"
	"github.com/weirdo0314/tiny-tiktok/pkg/bound"
	"github.com/weirdo0314/tiny-tiktok/pkg/middleware"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	config.Init()
	dao.Init()
	rpc.Init()
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry(strings.Split(config.Service.EtcdAddress, ";"))
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", config.Service.IP, config.Service.Port))
	if err != nil {
		panic(err)
	}

	klog.Info(config.Service)
	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Service.Name}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                       // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),                                             // registry
	)
	if err = svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
