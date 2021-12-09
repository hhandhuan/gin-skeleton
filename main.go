package main

import (
	"github.com/hhandhuan/gin-skeleton/configs"
	"github.com/hhandhuan/gin-skeleton/database"
	"github.com/hhandhuan/gin-skeleton/server"
)

func main() {
	// 初始化配置
	configs.ConfigInit()
	// 初始数据库
	database.MysqlInit()
	// 启动服务
	server.Run()
}
