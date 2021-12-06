package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/app/api"
	"github.com/hhandhuan/gin-skeleton/app/utils/response"
	"net/http"
)

func Register(e *gin.Engine) *gin.Engine {
	e.GET("ping", func(ctx *gin.Context) {
		response.Data(ctx, 0, "ok", nil)
	})
	e.GET("echo", api.Hello.Echo)
	e.NoRoute(func(ctx *gin.Context) {
		response.Data(ctx, http.StatusNotFound, "not found", nil)
	})
	return e
}
