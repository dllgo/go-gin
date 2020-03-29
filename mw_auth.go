package ginplus

import (
	"errors"

	jwtplus "github.com/dllgo/go-jwt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 用户授权中间件
func AuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if t := GetToken(c); t != "" {
			_, err := jwtplus.ParseToken(t)
			if err != nil {
				ResError(c, errors.New("无效的token"))
				return
			}
		}
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
	}
}
