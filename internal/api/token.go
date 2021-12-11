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
// @Success 200 object response.Result
// @Router /api/auth/token [post]
func (*auth) Token(ctx *gin.Context) {
	var request apiRequest.CreateAuthTokenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Error(ctx, errors.NewError(errors.ParamCode, err.Error()))
	}
	err, token := service.AuthService.CreateToken(&request)
	if err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, gin.H{"token": fmt.Sprintf("Bearer %s", token)})
	}
}

// Refresh 刷新令牌
// @Summary 刷新令牌
// @Schemes
// @Description 刷新令牌
// @Tags 授权管理
// @Accept json
// @Produce json
// @Success 200 object response.Result
// @Router /api/auth/token [post]
func (*auth) Refresh(ctx *gin.Context) {
	log.Println("test")
}
