package zlog

import (
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

// 使用全局的zap logger
var gLogger *zap.SugaredLogger

func Init(dir string, level zapcore.Level) {

	if gLogger == nil {

		core := zapcore.NewCore(getEncoder(false), getWriter(dir, level), level)

		logger := zap.New(core, zap.AddCaller())

		gLogger = logger.Sugar()

	}

}

// Flush 刷新日志,这里没有校验 gLogger 是否是 nil
func Flush() {
	gLogger.Sync()
}

// D 刷新日志,这里没有校验 gLogger 是否是 nil
func D(args ...interface{}) {
	gLogger.Debug(args...)
}

// I 刷新日志,这里没有校验 gLogger 是否是 nil
func I(args ...interface{}) {
	gLogger.Info(args...)
}

// E 刷新日志,这里没有校验 gLogger 是否是 nil
func E(args ...interface{}) {
	gLogger.Error(args...)
}

// W 刷新日志,这里没有校验 gLogger 是否是 nil
func W(args ...interface{}) {
	gLogger.Warn(args...)
}

// F 刷新日志,这里没有校验 gLogger 是否是 nil
func F(args ...interface{}) {
	gLogger.Fatal(args...)
}

func getEncoder(json bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if json {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriter(logDir string, lv Level) zapcore.WriteSyncer {

	fileName := filepath.Join(logDir, "cxbooks."+lv.CapitalString()+".log")

	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
