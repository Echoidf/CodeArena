package conf

import (
	"CodeArena/consts"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func SetLogger() {
	var zapOptions []zap.Option
	host, _ := os.Hostname()
	zapOptions = append(zapOptions,
		zap.AddCaller(),
		//zap.AddStacktrace(zap.ErrorLevel),
		zap.Fields(zap.String("hostname", host), zap.Int("pid", os.Getpid())))

	encodingConfig := zap.NewDevelopmentEncoderConfig()
	if SafeGetViperString(consts.ServerMode) != consts.DEV {
		encodingConfig = zap.NewProductionEncoderConfig()
	}

	//encodingConfig.ConsoleSeparator = " "
	encodingConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodingConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)

	var writeSyncers []zapcore.WriteSyncer

	writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout), getLogWriter())

	// 日志级别
	level, _ := zapcore.ParseLevel(V.GetString(consts.LogLevel))
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encodingConfig),
		zapcore.NewMultiWriteSyncer(writeSyncers...),
		level,
	)
	// 替换全局的日志记录器
	zap.ReplaceGlobals(zap.New(consoleCore, zapOptions...))
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   V.GetString("server.logPath"),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
