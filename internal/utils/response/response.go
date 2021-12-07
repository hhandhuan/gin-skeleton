package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"` // 业务码
	Msg  string      `json:"msg"`  // 业务消息
	Data interface{} `json:"data"` // 业务消息体
}

// Data 返回消息
func Data(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Error 错误返回
func Error(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
