package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hhandhuan/gin-skeleton/database"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"github.com/hhandhuan/gin-skeleton/internal/model"
	apiRequest "github.com/hhandhuan/gin-skeleton/internal/request"
)

var insUser = &userService{}

type userService struct{}

func User() *userService {
	return insUser
}

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

// EditUserInfo 编辑用户基础信息
func (u *userService) EditUserInfo(ctx *gin.Context, request *apiRequest.EditUserRequest) *errors.Error {
	err, user := u.GetLoggedUser(ctx)
	if err != nil {
		return err
	}
	database.Mysql.Table("users").Where("id = ?", user.ID).Updates(request)
	return nil
}

// GetUserByID 获取用户根据ID
func (*userService) GetUserByID(uid interface{}) (*errors.Error, *model.User) {
	var user model.User
	database.Mysql.Table("users").Where("id = ?", uid).First(&user)
	if user.ID <= 0 {
		return errors.NewError(errors.CommonCode, "user does not exist"), nil
	}
	return nil, &user
}
