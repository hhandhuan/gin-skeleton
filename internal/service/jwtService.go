package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hhandhuan/gin-skeleton/configs"
	"github.com/hhandhuan/gin-skeleton/database"
	"github.com/hhandhuan/gin-skeleton/internal/errors"
	"github.com/hhandhuan/gin-skeleton/internal/utils"
)

var JwtService = &jwtService{}

type jwtService struct{}

type Claims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

// MakeToken 生成令牌
func (*jwtService) MakeToken(uid int64) (string, error) {
	ttl := time.Hour * 24 * time.Duration(configs.Conf.Jwt.Ttl)
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Issuer:    "gin-skeleton",
		},
	})
	return tokenClaim.SignedString([]byte(configs.Conf.Jwt.Secret))
}

// ParseToken 解析令牌
func (*jwtService) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Conf.Jwt.Secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// MakeBlackListKey
func (*jwtService) MakeBlackListKey(token string) string {
	return fmt.Sprintf("jwt_black_list:%s", utils.Md5(token))
}

// InBlackList
func (a *jwtService) InBlackList(token string) bool {
	value, err := database.Redis.Get(context.Background(), a.MakeBlackListKey(token)).Result()
	if value == "" || err != nil {
		return false
	}
	timeValue, _ := strconv.ParseInt(value, 10, 64)
	if time.Now().Unix()-timeValue < 10 {
		return false
	}
	return true
}

// JoinBlackList
func (a *jwtService) JoinBlackList(token string) *errors.Error {
	claims, _ := a.ParseToken(token)
	currTime := time.Now().Unix()
	expired := time.Duration(claims.ExpiresAt-currTime) * time.Second
	err := database.Redis.SetNX(context.Background(), a.MakeBlackListKey(token), currTime, expired).Err()
	if err != nil {
		return errors.NewError(errors.CommonCode, err)
	}
	return nil
}
