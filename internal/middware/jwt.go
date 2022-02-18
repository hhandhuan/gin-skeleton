package middware

import (
	"time"

	"github.com/hhandhuan/gin-skeleton/internal/service"
	"github.com/hhandhuan/gin-skeleton/internal/utils"

	"github.com/hhandhuan/gin-skeleton/internal/errors"

	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/internal/utils/response"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err, token := utils.ParseTokenByHeader(ctx)
		if err != nil {
			response.Error(ctx, errors.NewError(errors.UNAuthCode, "invalid token"))
			ctx.Abort()
			return
		}
		claim, err := service.Jwt().ParseToken(token)
		if err != nil {
			response.Error(ctx, errors.NewError(errors.UNAuthCode, "invalid token"))
			ctx.Abort()
			return
		}
		// 提前五分钟过期
		if time.Now().Unix()-300 > claim.ExpiresAt {
			response.Error(ctx, errors.NewError(errors.UNAuthCode, "token has expired"))
			ctx.Abort()
			return
		}
		if service.Jwt().InBlackList(token) {
			response.Error(ctx, errors.NewError(errors.UNAuthCode, "invalid token"))
			ctx.Abort()
			return
		}
		ctx.Set("uid", claim.Uid)
		ctx.Next()
	}
}
