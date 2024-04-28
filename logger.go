package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap *zap.Logger
}

func New(level Level) (Logger, error) {
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(level)),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(zap.AddCallerSkip(skip))
	if err != nil {
		return Logger{}, errors.Wrap(err, "failed to build logger")
	}

	return Logger{
		logger,
	}, nil
}

func (l Logger) Sync() error {
	return l.zap.Sync()
}

func (l Logger) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, Fields(fields).zap()...)
}

func (l Logger) Info(msg string, fields ...Field) {
	l.zap.Info(msg, Fields(fields).zap()...)
}

func (l Logger) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, Fields(fields).zap()...)
}

func (l Logger) Error(msg string, fields ...Field) {
	l.zap.Error(msg, Fields(fields).zap()...)
}
