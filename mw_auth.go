package ginplus

import (
	"errors"

	jwtplus "github.com/dllgo/go-jwt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 用户授权中间件
func AuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := GetToken(c); t != "" {
			claims, err := jwtplus.ParseToken(t)
			if err != nil {
				ResError(c, errors.New("无效的token"))
				return
			}
			userID = claims.UserId
		}
		if userID != "" {
			c.Set(UserIDKey, userID)
		}
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		if userID == "" {
			ResError(c, errors.New("无权限"))
		}
	}
}
