package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/api"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

var (
	NotFoundErr       = "request not found"
	MethodNotFoundErr = "request method not found"
)

func Register(engine *gin.Engine) *gin.Engine {
	apiGroup := engine.Group("api")
	apiGroup.GET("user/details", api.User.Details)
	apiGroup.POST("user/create", api.User.Create)
	engine.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, 404, NotFoundErr)
	})
	engine.NoMethod(func(ctx *gin.Context) {
		response.Error(ctx, 405, MethodNotFoundErr)
	})
	return engine
}
