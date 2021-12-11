package service

import (
	"github.com/hhandhuan/gin-skeleton/database"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"github.com/hhandhuan/gin-skeleton/internal/model"
	"github.com/hhandhuan/gin-skeleton/internal/request"
	"github.com/hhandhuan/gin-skeleton/internal/utils"
	"github.com/hhandhuan/gin-skeleton/internal/utils/jwt"
)

var AuthService = &authService{}

type authService struct{}

// CreateToken CreateToken
func (*authService) CreateToken(request *request.CreateAuthTokenRequest) (*errors.Error, string) {
	var user model.User
	database.Mysql.Table("users").Where("username = ?", request.Username).First(&user)
	if user.ID <= 0 {
		return errors.NewError(errors.CommonCode, "user does not exist"), ""
	}
	if user.Password != utils.Md5(request.Password) {
		return errors.NewError(errors.CommonCode, "incorrect username or password"), ""
	}
	token, err := jwt.MakeToken(user.ID)
	if err != nil {
		return errors.NewError(errors.CommonCode, err), ""
	}
	return nil, token
}
