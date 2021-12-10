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
		apiGroup.POST("token/create", api.Auth.Token)
		apiGroup.Use(middware.JwtAuth())
		{
			apiGroup.POST("token/refresh", api.Auth.Token)
			apiGroup.GET("user/details", api.User.Details)
			apiGroup.POST("user/create", api.User.Create)
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
