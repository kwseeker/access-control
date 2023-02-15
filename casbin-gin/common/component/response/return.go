package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回值的统一定义

var Default = &response{}

// Error 失败数据处理
func Error(c *gin.Context, code int, err error, msg string) {
	res := Default.Clone()
	if err != nil {
		res.SetMsg(err.Error())
	}
	if msg != "" {
		res.SetMsg(msg)
	}
	//res.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	res.SetCode(int32(code))
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	res := Default.Clone()
	res.SetData(data)
	res.SetSuccess(true)
	if msg != "" {
		res.SetMsg(msg)
	}
	//res.SetTraceID(pkg.GenerateMsgIDFromContext(c))
	res.SetCode(http.StatusOK)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
