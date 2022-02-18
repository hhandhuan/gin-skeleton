package main

import (
	_ "github.com/hhandhuan/gin-skeleton/configs"
	_ "github.com/hhandhuan/gin-skeleton/database"
	_ "github.com/hhandhuan/gin-skeleton/logger"
	"github.com/hhandhuan/gin-skeleton/server"
)

// Run server run
// @title gin-skeleton
// @version 1.0
// @description gin-skeleton 示例项目
// @host 127.0.0.1:8081
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BashPth /
func main() {
	server.Run()
}
