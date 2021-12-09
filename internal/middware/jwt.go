package middware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		tokens := strings.Split(token, " ")
		if len(token) <= 0 || tokens[0] != "Bearer" || len(tokens[1]) <= 0 {
			response.Error(ctx, 4001, "token invlid")
			ctx.Abort()
			return
		}
	}
}
