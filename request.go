package ginplus

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/dllgo/go-utils"

	"github.com/gin-gonic/gin"
)

// 定义上下文中的键
const (
	prefix = "doudou"
	// UserIDKey 存储上下文中的键(用户ID)
	UserIDKey = prefix + "/user_id"
	// TraceIDKey 存储上下文中的键(跟踪ID)
	TraceIDKey = prefix + "/trace_id"
	// ResBodyKey 存储上下文中的键(响应Body数据)
	ResBodyKey = prefix + "/res_body"
)

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.New("无效的请求参数")
	}
	return nil
}

// GetUserID 获取用户ID
func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

// SetUserID 设定用户ID
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}

// GetToken 获取用户令牌
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetBasicToken 获取basic认证信息
func GetBasicToken(c *gin.Context) (string, string, error) {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Basic "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	credential, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", "", err
	}
	userAndPassword := strings.Split(string(credential), ":")
	return userAndPassword[0], userAndPassword[1], nil
}

// GetPageIndex 获取分页的页索引
func GetPageIndex(c *gin.Context) int {
	defaultVal := 1
	if v := c.Query("pageIndex"); v != "" {
		if iv := utils.S(v).DefaultInt(defaultVal); iv > 0 {
			return iv
		}
	}
	return defaultVal
}

// GetPageSize 获取分页的页大小(最大50)
func GetPageSize(c *gin.Context) int {
	defaultVal := 10
	if v := c.Query("pageSize"); v != "" {
		if iv := utils.S(v).DefaultInt(defaultVal); iv > 0 {
			if iv > 50 {
				iv = 50
			}
			return iv
		}
	}
	return defaultVal
}

// GetPaginationParam 获取分页查询参数
func GetPaginationParam(c *gin.Context) *HTTPPagination {
	return &HTTPPagination{
		PageIndex: GetPageIndex(c),
		PageSize:  GetPageSize(c),
	}
}
