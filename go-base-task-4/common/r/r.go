package r

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
响应格式统一
*/
type Response struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

// 失败数据处理
func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	if err != nil {
		res.Msg = err.Error()
	}
	if msg != "" {
		res.Msg = msg
	}
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnError(code))
}

func ErrMsg(c *gin.Context, msg string) {
	var res Response
	if msg != "" {
		res.Msg = msg
	}
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnError(http.StatusBadRequest))
}

// 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnOK())
}

// 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res Page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}
