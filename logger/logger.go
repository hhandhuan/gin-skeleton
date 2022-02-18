package logger

import (
	"github.com/hhandhuan/gin-skeleton/configs"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"path"
)

var I = logrus.New()

func init() {
	// 设置输出
	I.SetOutput(&lumberjack.Logger{
		Filename:   path.Join(configs.Conf.Logger.Filename), // 日志文件位置
		MaxSize:    configs.Conf.Logger.Maxsize,             // 单文件最大容量,单位是MB
		MaxBackups: configs.Conf.Logger.MaxBackups,          // 最大保留过期文件个数
		MaxAge:     configs.Conf.Logger.Maxage,              // 保留过期文件的最大时间间隔,单位是天
		Compress:   true,                                    // 是否需要压缩滚动日志, 使用的 gzip 压缩
	})
	// 设置日志输出格式
	I.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 设置日志记录级别
	I.SetLevel(logrus.Level(configs.Conf.Logger.Level))
}
