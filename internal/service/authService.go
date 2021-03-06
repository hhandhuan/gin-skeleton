package service

import (
	"fmt"
	"github.com/hhandhuan/gin-skeleton/database"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"github.com/hhandhuan/gin-skeleton/internal/model"
	"github.com/hhandhuan/gin-skeleton/internal/request"
	"github.com/hhandhuan/gin-skeleton/internal/utils"
)

var insAuth = &authService{}

type authService struct{}

func Auth() *authService {
	return insAuth
}

// CreateToken CreateToken
func (*authService) CreateToken(request *request.CreateAuthTokenRequest) (*errors.Error, string) {
	var user model.User
	database.Mysql.Table("users").Where("email = ?", request.Email).First(&user)
	if user.ID <= 0 {
		return errors.NewError(errors.CommonCode, "user does not exist"), ""
	}
	if user.Password != utils.Md5(request.Password) {
		return errors.NewError(errors.CommonCode, "incorrect username or password"), ""
	}
	token, err := jnsJwt.MakeToken(user.ID)
	if err != nil {
		return errors.NewError(errors.CommonCode, err), ""
	}
	return nil, fmt.Sprintf("Bearer %s", token)
}
