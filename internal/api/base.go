package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

var Base = &base{}

type base struct{}

// Ping 服务健康
// @Summary 服务健康
// @Schemes
// @Description 服务健康
// @Tags 基础管理
// @Accept json
// @Produce json
// @Success 200 object response.Result
// @Router /api/ping [get]
func (*base) Ping(ctx *gin.Context) {
	response.Data(ctx, "pong")
}
