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
		dataSource
	}
	server struct {
		Port string // Port 服务接口
		Name string // Name 服务名
		IP   string // IP 部署地址
	}
	// app 应用配置
	app struct {
		DefaultPageSize       uint32        // DefaultPageSize 默认分页大小
		MaxPageSize           uint32        // MaxPageSize 每页最大大小
		DefaultContextTimeout time.Duration // DefaultContextTimeout 默认上下文超时时间
		EtcdAddress           string        // EtcdAddress 注册中心
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
	// dataSource 数据库结构体
	dataSource struct {
		// MySQL Mysql配置
		MySQL struct {
			Host        string // Host 主机地址
			UserName    string // UserName 用户名称
			Password    string // Password 用户密码
			DBType      string // DBType 数据库类型
			DBName      string // DBName 数据库名称
			MaxOpenConn int    // MaxOpenConn 最大打开连接
			MaxIdleConn int    // MaxIdleConn 最大闲置连接
			Charset     string // Charset 字符集
			ParseTime   string // ParseTime 是否将数据库时间类型转换成Go时间类型
			TimeLocal   string // TimeLocal 时间区域
		}
		// Redis redis配置
		Redis struct {
			Addr     string // Addr redis服务地址
			Password string // Password redis密码
			DbIndex  int    // DbIndex 数据库下标
		}
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

	err = vp.UnmarshalKey("dataSource", &Service.dataSource)
	if err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}

	err = vp.UnmarshalKey("server", &Service.server)
	if err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}
}
