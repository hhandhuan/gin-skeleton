package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/api"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"github.com/hhandhuan/gin-skeleton/internal/middware"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

func Register(engine *gin.Engine) *gin.Engine {
	engine.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, errors.NewError(errors.CommonCode, "path not found"))
	})
	engine.NoMethod(func(ctx *gin.Context) {
		response.Error(ctx, errors.NewError(errors.CommonCode, "method not found"))
	})
	group := engine.Group("api")
	group.GET("ping", api.Base.Ping)
	group.POST("auth/token", api.Auth.Token)
	group.Use(middware.JwtAuth())
	{
		group.GET("auth/user", api.Auth.Logged)
		group.POST("auth/logout", api.Auth.Logout)
		group.POST("user/edit", api.User.Edit)
		group.GET("user/:uid", api.User.Show)
	}
	return engine
}
