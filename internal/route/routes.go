package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/api"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"github.com/hhandhuan/gin-skeleton/internal/middware"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

func Register(engine *gin.Engine) *gin.Engine {
	apiGroup := engine.Group("api")
	{
		apiGroup.POST("auth/token", api.Auth.Token)
		apiGroup.Use(middware.JwtAuth())
		{
			apiGroup.GET("auth/user", api.Auth.Logged)
			apiGroup.POST("token/refresh", api.Auth.Refresh)
		}
	}
	engine.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, errors.NewError(errors.CommonCode, "not found"))
	})
	engine.NoMethod(func(ctx *gin.Context) {
		response.Error(ctx, errors.NewError(errors.CommonCode, "method not found"))
	})
	return engine
}
