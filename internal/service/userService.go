package service

import (
	"github.com/hhandhuan/gin-skeleton/database"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"github.com/hhandhuan/gin-skeleton/internal/model"
)

var (
	UserService = &userService{}
)

type userService struct{}

// GetUsers 获取用户列表
func (*userService) GetUsers() (users []*model.User) {
	database.Mysql.Table("users").Find(&users)
	return
}

func (*userService) GetUserByUsername(username string) (*errors.Error, *model.User) {
	var user *model.User
	database.Mysql.Table("users").Where("username = ?", username).First(user)
	if user.ID <= 0 {
		return errors.NewError(errors.CommonCode, "用户不存在"), nil
	}
	return nil, user
}
