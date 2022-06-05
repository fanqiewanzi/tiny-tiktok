package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/app"
	"github.com/weirdo0314/tiny-tiktok/cmd/api/pkg/util/auth"
	"github.com/weirdo0314/tiny-tiktok/pkg/errorx"
)

// JWT 验证JWT令牌中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.GetQuery("token")
		if !ok {
			token, ok = c.GetPostForm("token")
			//若token为空则直接返回,不为空则解析token
			if !ok {
				app.WriteResponse(c, errorx.AuthErr, nil)
				c.Abort()
				return
			}
		}

		claims, err := auth.ParseToken(token)
		if err != nil {
			app.WriteResponse(c, errorx.AuthCheckTokenErr, nil)
			c.Abort()
			return
		}

		curTime := time.Now()
		// TOKEN过期
		if !claims.VerifyExpiresAt(curTime, false) {
			app.WriteResponse(c, errorx.AuthCheckTokenTimeOutErr, nil)
			c.Abort()
			return
		}

		if !claims.VerifyIssuedAt(curTime, false) {
			app.WriteResponse(c, errorx.AuthErr, nil)
			c.Abort()
			return
		}

		if !claims.VerifyNotBefore(curTime, false) {
			app.WriteResponse(c, errorx.AuthErr, nil)
			c.Abort()
			return
		}

		// 将包含用户信息的载荷传递下去
		c.Set("claims", claims)
	}
}
