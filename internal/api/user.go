package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/service"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

var User = &user{}

type user struct{}

// @BasePath /api/user

// Details godoc
// @Summary 用户详情
// @Schemes
// @Description 用户详情
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 object []model.User
// @Router /api/user/details [get]
func (u *user) Details(ctx *gin.Context) {
	response.Data(ctx, service.UserService.GetUsers())
}

// Create godoc
// @Summary 用户创建
// @Schemes
// @Description 用户创建
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 object response.Result
// @Router /api/user/create [post]
func (u *user) Create(ctx *gin.Context) {
	response.Data(ctx, &response.Result{})
}
