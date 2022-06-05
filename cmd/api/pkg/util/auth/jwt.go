package auth

import (
	"time"

	"github.com/weirdo0314/tiny-tiktok/cmd/api/config"

	"github.com/golang-jwt/jwt/v4"
)

// Claims 自定义载荷结构，包含了一个标准的JWT载荷
type Claims struct {
	// ID 用户ID
	ID uint64 `json:"id"`
	// jwt标准载荷
	jwt.RegisteredClaims
}

// ParseToken 解析token字符串
func ParseToken(token string) (*Claims, error) {
	p := jwt.NewParser(jwt.WithoutClaimsValidation())
	//用密匙解析出token声明
	tokenClaims, err := p.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Service.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	// token声明解析完成后再次解析出自定义的声明并检查token声明是否有效，有效则返回自定义声明
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}

// CreateToken 创建Token
func CreateToken(id uint64) (string, error) {
	//通过现在时间得到token过期时间
	curTime := time.Now()
	expTime := curTime.Add(config.Service.Timeout * time.Minute)
	//初始化自定义的载荷
	claims := Claims{id, jwt.RegisteredClaims{ExpiresAt: &jwt.NumericDate{Time: expTime},
		Issuer: config.Service.Issuer, IssuedAt: &jwt.NumericDate{Time: curTime}}}
	//初始化tokenClaim数据结构
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaim.SignedString([]byte(config.Service.Secret))
	return token, err
}
