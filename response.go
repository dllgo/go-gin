package ginplus

import (
	"net/http"

	"github.com/dllgo/go-utils"

	"github.com/gin-gonic/gin"
)

const (
	MSG_OK       = 0  // 成功
	MSG_ERR      = -1 // 错误
	MSG_REDIRECT = -2 // 重定向
	MSG_AUTH     = -3 // 登录认证
)

// HTTPItem HTTP响应错误项
type HTTPError struct {
	Code    int    `json:"code" swaggo:"true,错误码"`
	Message string `json:"message" swaggo:"true,错误信息"`
}

// HTTPItem HTTP响应项
type HTTPItem struct {
	Code    int         `json:"code" swaggo:"true,状态码"`
	Message string      `json:"message" swaggo:"true,状态信息"`
	Data    interface{} `json:"data" swaggo:"true,状态数据"`
}

// HTTPStatus HTTP响应状态
type HTTPStatus struct {
	Status string `json:"status" swaggo:"true,状态(OK)"`
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	List       interface{}     `json:"list"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total     int `json:"total"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

// ResList 响应列表数据
func ResList(c *gin.Context, v interface{}, total int) {
	list := HTTPList{
		List: v,
		Pagination: &HTTPPagination{
			Total:     total,
			PageIndex: GetPageIndex(c),
			PageSize:  GetPageSize(c),
		},
	}
	ResSuccess(c, list)
}

// ResOK 响应OK
func ResOK(c *gin.Context, v interface{}) {
	ResSuccess(c, v)
}

// ResError 响应错误
func ResError(c *gin.Context, err error) {
	errItem := HTTPError{
		Code:    100,
		Message: err.Error(),
	}
	ResJSON(c, errItem)
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	data := HTTPItem{
		Code:    200,
		Message: "ok",
		Data:    v,
	}
	ResJSON(c, data)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, v interface{}) {
	buf, err := utils.JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(http.StatusOK, "application/json; charset=utf-8", buf)
	c.Abort()
}
