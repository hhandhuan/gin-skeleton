package main

import (
	"github.com/hhandhuan/gin-skeleton/configs"
	"github.com/hhandhuan/gin-skeleton/database"
	_ "github.com/hhandhuan/gin-skeleton/docs"
	"github.com/hhandhuan/gin-skeleton/server"
)


// @title gin-skeleton
// @version 1.0
// @description gin-skeleton 示例项目
// @host 127.0.0.1:8080
func main() {
	configs.InitConfig()
	database.MysqlInit()
	server.Run()
}