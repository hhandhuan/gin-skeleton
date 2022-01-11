package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormDefaultLogger "gorm.io/gorm/logger"

	"github.com/hhandhuan/gin-skeleton/configs"
)

var Mysql *gorm.DB

// MysqlInit 初始化数据库
func MysqlInit() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       configs.Conf.Database.Link, // DSN data source name
		DefaultStringSize:         255,                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: gormDefaultLogger.Default.LogMode(gormDefaultLogger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	// 设置可以重用连接的最长时间
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	Mysql = db
}
