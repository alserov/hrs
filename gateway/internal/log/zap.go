package log

import (
	"github.com/alserov/hrs/gateway/internal/config"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

var log Logger

func GetLogger() Logger {
	return log
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

	log = Logger{l}

	return log
}
