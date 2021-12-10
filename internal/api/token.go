package api

import (
	"github.com/hhandhuan/gin-skeleton/internal/utils/jwt"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
	"log"

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
	token, err := jwt.MakeToken(11)
	if err != nil {
		response.Error(ctx, 5001, "make token error")
		return
	}
	response.Data(ctx, token)
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