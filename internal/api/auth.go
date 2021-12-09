package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

var Auth = &auth{}

type auth struct{}

func (*auth) Token(ctx *gin.Context) {
	log.Println("test")
}

func (*auth) Refresh(ctx *gin.Context) {
	log.Println("test")
}
