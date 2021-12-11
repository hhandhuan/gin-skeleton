package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	apiRequest "github.com/hhandhuan/gin-skeleton/internal/request"
	"github.com/hhandhuan/gin-skeleton/internal/service"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

var User = &user{}

type user struct{}

// Edit 修改基础信息
// @Summary 修改基础信息
// @Schemes
// @Description 修改基础信息
// @Tags 授权管理
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
	if err := service.UserService.EditUserInfo(ctx, &request); err != nil {
		response.Error(ctx, err)
		return
	} else {
		response.Data(ctx, nil)
	}
}
