package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
	"github.com/hhandhuan/gin-skeleton/logger"
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
	for i := 0; i < 1000; i++ {
		logger.I.WithField("username", "erici")
	}
	response.Data(ctx, "pong")
}
