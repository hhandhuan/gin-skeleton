package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/app/utils/response"
)

var (
	Hello = &hello{}
)

type hello struct{}

func (h *hello) Echo(ctx *gin.Context) {
	response.Data(ctx, 0, "ok", nil)
}
