package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/weirdo0314/tiny-tiktok/cmd/video/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	db  *gorm.DB
	rdb *redis.Client
)

func Init() {
	InitMysql()
	InitRDB()
}

// InitMysql init mysql
func InitMysql() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: config.Service.MySQL.DBType,
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
			config.Service.MySQL.UserName, config.Service.MySQL.Password, config.Service.MySQL.Host, config.Service.MySQL.DBName,
			config.Service.MySQL.Charset, config.Service.MySQL.ParseTime, config.Service.MySQL.TimeLocal),
		DefaultStringSize: 256}),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(255) //设置最大连接数
	sqlDB.SetMaxOpenConns(100) //设置最大的空闲连接数
	sqlDB.SetConnMaxIdleTime(10 * time.Second)

	if err = db.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
}

// InitRDB init redis
func InitRDB() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Service.Redis.Addr,
		Password: config.Service.Redis.Password,
		DB:       config.Service.Redis.DbIndex,
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("redis 连接失败")
	}
}
