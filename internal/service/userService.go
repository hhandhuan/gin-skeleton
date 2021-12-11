package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/database"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"github.com/hhandhuan/gin-skeleton/internal/model"
)


var UserService = &userService{}

type userService struct{}

// GetLoggedUser 获取当前登陆的用户
func (*userService) GetLoggedUser(ctx *gin.Context) (*errors.Error, *model.User) {
	var user model.User
	uid, ok := ctx.Get("uid")
	if !ok {
		return errors.NewError(errors.CommonCode, "user does not exist"), nil
	}
	database.Mysql.Table("users").Where("id = ?", uid).First(&user)
	if user.ID <= 0 {
		return errors.NewError(errors.CommonCode, "user does not exist"), nil
	}
	return nil, &user
}
