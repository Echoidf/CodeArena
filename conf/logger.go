package conf

import (
	"CodeArena/consts"
	"CodeArena/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func SetLogger() {
	logger := zap.NewExample()
	defer logger.Sync()

	var zapOptions []zap.Option
	host, _ := os.Hostname()
	zapOptions = append(zapOptions,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
		zap.Fields(zap.String("hostname", host), zap.Int("pid", os.Getpid())))

	encodingConfig := zap.NewDevelopmentEncoderConfig()
	if utils.SafeGetString(consts.ServerMode) != consts.DEV {
		encodingConfig = zap.NewProductionEncoderConfig()
	}

	encodingConfig.ConsoleSeparator = " "
	encodingConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encodingConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)

	var writeSyncers []zapcore.WriteSyncer

	logPath := V.GetString("server.logPath")
	var logFile *os.File
	if utils.NotExistFile(logPath) {
		logFile, _ = utils.CreateFile(logPath)
	} else {
		logFile, _ = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	}

	writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout), zapcore.AddSync(logFile))

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encodingConfig),
		zapcore.NewMultiWriteSyncer(writeSyncers...),
		nil,
	)
	// 替换全局的日志记录器
	zap.ReplaceGlobals(zap.New(consoleCore, zapOptions...))
}
