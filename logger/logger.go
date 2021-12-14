package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/hhandhuan/gin-skeleton/configs"
	"github.com/natefinch/lumberjack"
)

var Log *zap.Logger

// LoggerInit 初始化Logger
func LoggerInit() {
	writeSyncer := getLogWriter(
		configs.Conf.Logger.Filename,
		configs.Conf.Logger.Maxsize,
		configs.Conf.Logger.MaxBackups,
		configs.Conf.Logger.Maxage,
	)
	encoder := getEncoder()
	var level = new(zapcore.Level)
	err := level.UnmarshalText([]byte(configs.Conf.Logger.Level))
	if err != nil {
		log.Fatal(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)

	Log = zap.New(core, zap.AddCaller())
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(Log)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
