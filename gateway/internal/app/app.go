package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/alserov/hrs/gateway/internal/adapters"
	"github.com/alserov/hrs/gateway/internal/clients"
	"github.com/alserov/hrs/gateway/internal/config"
	"github.com/alserov/hrs/gateway/internal/log"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func MustStart(cfg *config.Config) {
	l := log.NewLogger(cfg.Env)

	defer func() {
		if err := recover(); err != nil {
			l.Error("recovery ❌", zap.Any("error", err))
		}
	}()

	cls := clients.NewClients()

	ctrl := adapters.NewController(cls)

	router := echo.New()
	adapters.Setup(router, ctrl, l)

	l.Info("layers are set up ✔")

	shutdown(func() {
		l.Info("server is running ✔", zap.Int("port", cfg.Port))
		run(cfg.Port, router)
	}, router)

	l.Info("server stopped ✔")
}

func run(port int, s *echo.Echo) {
	if err := s.Start(fmt.Sprintf(":%d", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("failed to start server: " + err.Error())
	}
}

func shutdown(fn func(), s *echo.Echo) {
	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	go fn()

	<-chStop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		panic("failed to shutdown server: " + err.Error())
	}
}
