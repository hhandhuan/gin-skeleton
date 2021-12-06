package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/api"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
	"net/http"
)

var (
	NotFoundErr       = "request not found"
	MethodNotFoundErr = "request method not found"
)

func Register(e *gin.Engine) *gin.Engine {
	e.GET("user/details", api.User.Details)
	e.POST("user/create", api.User.Create)
	e.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, http.StatusNotFound, NotFoundErr)
	})
	e.NoMethod(func(ctx *gin.Context) {
		response.Error(ctx, http.StatusNotFound, MethodNotFoundErr)
	})
	return e
}
