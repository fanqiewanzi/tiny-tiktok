package filename

import (
	"math/rand"
	"strings"
	"time"
)

// getRandstring 获取随机字母+数字组合字符串
func getRandstring(length int) string {
	if length < 1 {
		return ""
	}

	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

// RandFileName 随机文件名
func RandFileName() string {
	randStr := getRandstring(16)
	return randStr
}
