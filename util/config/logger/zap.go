package logger

import (
	"fmt"
	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

const (
	// logFormat
	//LOGFORMAT_JSON    = "json"
	//LOGFORMAT_CONSOLE = "console"

	// EncoderConfig
	TIME_KEY       = "time"
	LEVLE_KEY      = "level"
	NAME_KEY       = "logger"
	CALLER_KEY     = "caller"
	MESSAGE_KEY    = "msg"
	STACKTRACE_KEY = "stacktrace"

	MAX_SIZE               = 1
	MAX_BACKUPS            = 5
	MAX_AGE                = 7
	DebugLevelStr   string = "debug"
	InfoLevelStr    string = "info"
	WarningLevelStr string = "warning"
	ErrorLevelStr   string = "error"
)

func Init(logLevel string, fileName string, useLogFile bool, dsn string) error {
	var level zapcore.Level
	switch strings.ToLower(logLevel) {
	case DebugLevelStr:
		level = zap.DebugLevel
	case InfoLevelStr:
		level = zap.InfoLevel
	case WarningLevelStr:
		level = zap.WarnLevel
	case ErrorLevelStr:
		level = zap.ErrorLevel
	default:
		return fmt.Errorf("unknown log level %s", logLevel)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        TIME_KEY,
		LevelKey:       LEVLE_KEY,
		NameKey:        NAME_KEY,
		CallerKey:      CALLER_KEY,
		MessageKey:     MESSAGE_KEY,
		StacktraceKey:  STACKTRACE_KEY,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewNopCore()

	if useLogFile {
		core = zapcore.NewCore(
			encoder,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&lumberjack.Logger{
				Filename:   fileName,
				MaxSize:    MAX_SIZE,
				MaxBackups: MAX_BACKUPS,
				MaxAge:     MAX_AGE,
				Compress:   true,
			})),
			zap.NewAtomicLevelAt(level),
		)
	} else {
		core = zapcore.NewCore(
			encoder,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			zap.NewAtomicLevelAt(level),
		)
	}

	caller := zap.AddCaller()
	development := zap.Development()
	logger := zap.New(core, caller, development)

	cfg := zapsentry.Configuration{
		Level: zapcore.ErrorLevel, //when to send message to sentry
		Tags: map[string]string{
			"component": "system",
		},
	}
	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(dsn))
	//in case of err it will return noop core. so we can safely attach it
	if err != nil {
		logger.Warn("failed to init zap", zap.Error(err))
	}
	//return zapsentry.AttachCoreToLogger(core, logger)

	zap.ReplaceGlobals(zapsentry.AttachCoreToLogger(core, logger))
	//zap.ReplaceGlobals(logger)

	zap.L().Info("zap logger and sentry configure is done")

	return nil
}
