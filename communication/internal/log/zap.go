package log

import (
	"context"
	"github.com/alserov/hrs/communication/internal/config"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

const (
	CtxLogger = "logger"
)

func WithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, CtxLogger, log)
}

func GetLogger(ctx context.Context) Logger {
	l, _ := ctx.Value(CtxLogger).(Logger)
	return l
}

func NewLogger(env string) Logger {
	var (
		l   *zap.Logger
		err error
	)

	switch env {
	case config.Local:
		l, err = zap.NewDevelopment()
	case config.Production:
		l, err = zap.NewProduction()
	default:
		panic("unknown env: " + env)
	}

	if err != nil {
		panic("failed to init logger")
	}

	return Logger{l}
}
