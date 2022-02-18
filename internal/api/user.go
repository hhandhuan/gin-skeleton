package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	apiRequest "github.com/hhandhuan/gin-skeleton/internal/request"
	"github.com/hhandhuan/gin-skeleton/internal/service"
	"github.com/hhandhuan/gin-skeleton/internal/utils"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

var User = &user{}

type user struct{}

// Edit 修改基础信息
// @Security BearerAuth
// @Summary 修改基础信息
// @Schemes
// @Description 修改基础信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param params body apiRequest.EditUserRequest true "用户名, 用户呢称"
// @Success 200 object response.Result
// @Router /api/user/edit [post]
func (*user) Edit(ctx *gin.Context) {
	var request apiRequest.EditUserRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.Error(ctx, errors.NewError(errors.ParamCode, err))
		return
	}
	if err := service.User().EditUserInfo(ctx, &request); err != nil {
		response.Error(ctx, err)
		return
	} else {
		response.Data(ctx, nil)
	}
}

// Show 获取用户详情
// @Security BearerAuth
// @Summary 获取用户详情
// @Schemes
// @Description 获取用户详情
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 object response.Result
// @Param uid path int true "user ID"
// @Router /api/user/{uid} [get]
func (*user) Show(ctx *gin.Context) {
	uid := utils.StrToInt(ctx.Param("uid"))
	if uid <= 0 {
		response.Error(ctx, errors.NewError(errors.ParamCode, "参数错误"))
		return
	}
	err, user := service.User().GetUserByID(uid)
	if err != nil {
		response.Error(ctx, err)
		return
	} else {
		response.Data(ctx, user)
	}
}
