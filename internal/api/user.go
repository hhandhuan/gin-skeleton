package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

var User = &user{}

type user struct{}

// @BasePath /user

// Details godoc
// @Summary 用户详情
// @Schemes
// @Description hello echo
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 object response.Result
// @Router /user/details [get]
func (u *user) Details(ctx *gin.Context) {
	response.Data(ctx, 0, "ok", nil)
}

// Create godoc
// @Summary 用户创建
// @Schemes
// @Description hello echo
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 object response.Result
// @Router /user/create [post]
func (u *user) Create(ctx *gin.Context) {
	response.Data(ctx, 0, "ok", nil)
}
