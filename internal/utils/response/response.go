package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Data(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": data})
}

func Error(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}
