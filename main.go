package main

import (
	"github.com/hhandhuan/gin-skeleton/configs"
	"github.com/hhandhuan/gin-skeleton/database"
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
	// 初始化配置
	configs.ConfigInit()
	// 初始数据库
	database.MysqlInit()
	// 启动服务
	server.Run()
}
