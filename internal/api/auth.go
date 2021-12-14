package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	apiRequest "github.com/hhandhuan/gin-skeleton/internal/request"
	"github.com/hhandhuan/gin-skeleton/internal/service"
	"github.com/hhandhuan/gin-skeleton/internal/utils"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

var Auth = &auth{}

type auth struct{}

// Token 创建令牌
// @Summary 令牌创建
// @Schemes
// @Description 令牌创建
// @Tags 授权管理
// @Accept json
// @Produce json
// @Param params body apiRequest.CreateAuthTokenRequest true "用户名, 密码"
// @Success 200 object response.Result
// @Router /api/auth/token [post]
func (*auth) Token(ctx *gin.Context) {
	var request apiRequest.CreateAuthTokenRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.Error(ctx, errors.NewError(errors.ParamCode, err))
		return
	}
	err, token := service.AuthService.CreateToken(&request)
	if err != nil {
		response.Error(ctx, err)
		return
	} else {
		result := gin.H{"token": fmt.Sprintf("Bearer %s", token)}
		response.Data(ctx, result)
	}
}

// Logged 获取用户信息
// @Security BearerAuth
// @Summary 获取用户信息
// @Schemes
// @Description 获取用户信息
// @Tags 授权管理
// @Accept json
// @Produce json
// @Success 200 object response.Result
// @Router /api/auth/user [get]
func (*auth) Logged(ctx *gin.Context) {
	if err, user := service.UserService.GetLoggedUser(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, user)
	}
}

// logout 退出登录
// @Security BearerAuth
// @Summary 退出登录
// @Schemes
// @Description 退出登录
// @Tags 授权管理
// @Accept json
// @Produce json
// @Success 200 object response.Result
// @Router /api/auth/logout [get]
func (*auth) Logout(ctx *gin.Context) {
	_, token := utils.ParseTokenByHeader(ctx)
	if err := service.JwtService.JoinBlackList(token); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, nil)
	}
}
