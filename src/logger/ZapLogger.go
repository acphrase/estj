package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// singleton 객체값(pointer)
var log *zap.Logger

func init() {
	config := setLoggerConfig()

	var err error
	log, err = config.Build(zap.AddCallerSkip(1)) // 설정 완료 후 logger build.
	if err != nil {
		panic(err)
	}
}

func setLoggerConfig() zap.Config {

	// TODO: profile 설정 진행 후 변경 예정.
	logOutputList := []string{
		"stdout",
		"/Volumes/Data/estj/src/estj.log",
	}
	logOutputListForLoggerError := []string{
		"stdout",
		"/Volumes/Data/estj/src/logger_error.log",
	}

	// Set logger config.
	var loggerConfig zap.Config = zap.NewProductionConfig()

	// Set logging level.
	// TODO: profile 설정 진행 후 변경 예정.
	loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

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
