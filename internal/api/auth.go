package api

import (
	"fmt"
	"log"

	"github.com/hhandhuan/gin-skeleton/internal/errors"
	apiRequest "github.com/hhandhuan/gin-skeleton/internal/request"
	"github.com/hhandhuan/gin-skeleton/internal/service"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"

	"github.com/gin-gonic/gin"
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
	if err := ctx.ShouldBindJSON(&request); err != nil {
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

// Refresh 刷新令牌
// @Security BearerAuth
// @Summary 刷新令牌
// @Schemes
// @Description 刷新令牌
// @Tags 授权管理
// @Accept json
// @Produce json
// @Success 200 object response.Result
// @Router /api/auth/refresh [post]
func (*auth) Refresh(ctx *gin.Context) {
	log.Println("test")
}
