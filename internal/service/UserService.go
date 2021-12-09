package service

import (
	"github.com/hhandhuan/gin-skeleton/database"
	"github.com/hhandhuan/gin-skeleton/internal/model"
)

var UserService = &userService{}

type userService struct{}

// GetUsers 获取用户列表
func (*userService) GetUsers() (users []*model.User) {
	database.Mysql.Table("users").Find(&users)
	return
}