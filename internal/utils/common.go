package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func ParseTokenByHeader(ctx *gin.Context) (error, string) {
	token := ctx.Request.Header.Get("Authorization")
	tokens := strings.Split(token, " ")
	if len(token) <= 0 || len(tokens[1]) <= 0 {
		return errors.New("parse token error"), ""
	}
	return nil, tokens[1]
}

func StrToInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}
