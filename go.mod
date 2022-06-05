module github.com/weirdo0314/tiny-tiktok

go 1.16

require (
	github.com/apache/thrift v0.15.0 // indirect
	github.com/cloudwego/kitex v0.3.1
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/kitex-contrib/registry-etcd v0.0.0-20220110034026-b1c94979cea3
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/qiniu/go-sdk/v7 v7.12.1
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/viper v1.11.0
	github.com/tidwall/gjson v1.12.1 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v3 v3.0.0-20220512140231-539c8e751b99 // indirect
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
	gorm.io/plugin/opentracing v0.0.0-20211220013347-7d2b2af23560
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
