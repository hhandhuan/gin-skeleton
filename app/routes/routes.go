package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/app/api"
	"github.com/hhandhuan/gin-skeleton/app/utils/response"
	"net/http"
)

var (
	NotFoundErr       = "not found"
	MethodNotFoundErr = "method not found"
)

func Register(e *gin.Engine) *gin.Engine {
	e.GET("ping", func(ctx *gin.Context) {
		response.Data(ctx, 0, "ok", nil)
	})
	e.GET("echo", api.Hello.Echo)
	e.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, http.StatusNotFound, NotFoundErr)
	})
	e.NoMethod(func(ctx *gin.Context) {
		response.Error(ctx, http.StatusNotFound, MethodNotFoundErr)
	})
	return e
}
