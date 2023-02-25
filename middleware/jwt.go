package middleware

import (
	"school_manager/pkg"
	"school_manager/util"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = pkg.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = pkg.ErrorAuthCheckTokenTimeout
			}
		}
		if code != pkg.SUCCESS {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    pkg.GetMsg(code),
				"data":   "可能是身份过期了，请重新登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
