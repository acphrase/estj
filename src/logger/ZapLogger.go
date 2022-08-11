package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// singleton 객체값(pointer)
var log *zap.Logger

func InitLogger(logLevel string, logOutputList []string, logOutputListForLoggerError []string) {
	if log == nil {
		config := initLoggerConfig(logLevel, logOutputList, logOutputListForLoggerError)

		var err error
		log, err = config.Build(zap.AddCallerSkip(1)) // 설정 완료 후 logger build.
		if err != nil {
			panic(err)
		}
	}
}

func initLoggerConfig(logLevel string, logOutputList []string, logOutputListForLoggerError []string) zap.Config {

	// Set logger config.
	var loggerConfig zap.Config = zap.NewProductionConfig()

	// Set logging level.
	switch logLevel {
	case "DPANIC", "dpanic":
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "DEBUG", "debug":
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "ERROR", "error":
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "FATAL", "fatal":
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	case "PANIC", "panic":
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "INFO", "info":
		fallthrough
	default:
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	// Set json encoder config.
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // timestamp를 보기 좋게 하기 위한 옵션.
	encoderConfig.StacktraceKey = ""                      // error가 발생 했을 때, stack push를 비어있도록 하는 설정.

	// Set json encoder config in logger config.
	loggerConfig.EncoderConfig = encoderConfig

	// Set output method setting.
	loggerConfig.OutputPaths = logOutputList
	// Set error output method setting.
	loggerConfig.ErrorOutputPaths = logOutputListForLoggerError

	return loggerConfig
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}

func DPanic(message string, fields ...zap.Field) {
	log.DPanic(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	log.Panic(message, fields...)
}
