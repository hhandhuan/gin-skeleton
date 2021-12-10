package middware

import (
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"github.com/hhandhuan/gin-skeleton/internal/utils/jwt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		tokens := strings.Split(token, " ")
		if len(token) <= 0 || tokens[0] != "Bearer" || len(tokens[1]) <= 0 {
			response.Error(ctx, errors.NewError(errors.AuthCode, "invalid token"))
			ctx.Abort()
			return
		}
		claim, err := jwt.ParseToken(tokens[1])
		if err != nil {
			response.Error(ctx, errors.NewError(errors.AuthCode, err.Error()))
			ctx.Abort()
			return
		}
		if time.Now().Unix() > claim.ExpiresAt {
			response.Error(ctx, errors.NewError(errors.AuthCode, "token has expired"))
			ctx.Abort()
			return
		}
		ctx.Set("uid", claim.Uid)
		ctx.Next()
	}
}
