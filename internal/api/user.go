package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

var (
	User = &user{}
)

type R struct {
	Code string      `json:"code" example:"0"`
	Msg  string      `json:"msg" example:"信息"`
	Data interface{} `json:"data"`
}

type user struct{}

// @BasePath /echo

// Details godoc
// @Summary 用户详情
// @Schemes
// @Description hello echo
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 object R
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
// @Success 200 object R
// @Router /user/create [post]
func (u *user) Create(ctx *gin.Context) {
	response.Data(ctx, 0, "ok", nil)
}
