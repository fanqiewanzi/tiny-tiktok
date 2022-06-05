package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

var Service Config

type (
	Config struct {
		server
		app
		jwt
		rpc
	}
	// server HTTP服务配置
	server struct {
		RunMode      string        // RunMode 运行模式
		HttpPort     string        // HttpPort 服务端口号
		ReadTimeout  time.Duration // ReadTimeout 读取超时时间
		WriteTimeout time.Duration // WriteTimeout 写入超时时间
	}
	// app 应用配置
	app struct {
		DefaultPageSize       uint32        // DefaultPageSize 默认分页大小
		MaxPageSize           uint32        // MaxPageSize 每页最大大小
		DefaultContextTimeout time.Duration // DefaultContextTimeout 默认上下文超时时间
		// 日志配置
		Log struct {
			SavePath   string // SavePath 日志保存地址
			FileName   string // FileName 日志文件名称
			FileExt    string // FileExt 日志文件扩展名
			MaxSize    int    // MaxSize 日志切割文件的最大大小
			MaxBackUps int    // MaxBackUps 保留旧文件的最大个数
			MaxAge     int    // MaxAges 保留旧文件的最大天数
		}
	}

	// jwt JWT验证配置
	jwt struct {
		Secret     string        // Secret JWT 密钥
		Issuer     string        // Issuer JWT 发行人
		Timeout    time.Duration // Timeout 过期时间
		MaxRefresh time.Duration // MaxRefresh 最大刷新时间
	}

	rpc struct {
		EtcdAddress      string
		UserServiceName  string `mapstructure:"user"`
		VideoServiceName string `mapstructure:"video"`
	}
)

func Init() {
	path := os.Getenv("CONFIG_PATH")
	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName("config")
	vp.AddConfigPath(path)
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("no such config file: %v\n", err)
		} else {
			log.Fatalf("read config error: %v\n", err)
		}
	}

	err := vp.UnmarshalKey("app", &Service.app)
	if err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}

	err = vp.UnmarshalKey("jwt", &Service.jwt)
	if err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}

	Service.Timeout *= time.Hour
	Service.MaxRefresh *= time.Hour

	err = vp.UnmarshalKey("rpc", &Service.rpc)
	if err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}

	err = vp.UnmarshalKey("server", &Service.server)
	if err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}

	Service.server.ReadTimeout *= time.Second
	Service.server.WriteTimeout *= time.Second
}
