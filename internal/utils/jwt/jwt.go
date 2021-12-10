package jwt

import (
	"github.com/hhandhuan/gin-skeleton/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

// MakeToken 生成令牌
func MakeToken(uid int64) (string, error) {
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
func ParseToken(token string) (*Claims, error) {
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
