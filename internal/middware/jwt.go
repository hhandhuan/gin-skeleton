package middware

import (
	"github.com/hhandhuan/gin-skeleton/internal/utils/jwt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

var (
	TokenInvalidErr = "invalid token"
	TokenExpiredErr = "token has expired"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		tokens := strings.Split(token, " ")
		if len(token) <= 0 || tokens[0] != "Bearer" || len(tokens[1]) <= 0 {
			response.Error(ctx, 4001, TokenInvalidErr)
			ctx.Abort()
			return
		}
		claim, err := jwt.ParseToken(tokens[1])
		if err != nil {
			response.Error(ctx, 4001, err.Error())
			ctx.Abort()
			return
		}
		if time.Now().Unix() > claim.ExpiresAt {
			response.Error(ctx, 4001, TokenExpiredErr)
			ctx.Abort()
			return
		}
		ctx.Set("uid", claim.Uid)
		ctx.Next()
	}
}
