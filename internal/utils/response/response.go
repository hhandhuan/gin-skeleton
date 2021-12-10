package response

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"net/http"
)

type Result struct {
	Code int         `json:"code"` // 业务码
	Msg  string      `json:"msg"`  // 业务消息
	Data interface{} `json:"data"` // 业务消息体
}

// Data 返回消息
func Data(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Result{Code: 0, Msg: "ok", Data: data})
}

// Error 错误返回
func Error(ctx *gin.Context, error *errors.Error) {
	ctx.JSON(http.StatusOK, &Result{Code: error.Code, Msg: error.Msg})
}
