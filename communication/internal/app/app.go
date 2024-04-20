package app

import (
	"fmt"
	"github.com/alserov/hrs/communication/internal/adapter/grpc"
	"github.com/alserov/hrs/communication/internal/config"
	"github.com/alserov/hrs/communication/internal/db/scylla"
	"github.com/alserov/hrs/communication/internal/log"
	"github.com/alserov/hrs/communication/internal/usecase"
	"go.uber.org/zap"
	gRPC "google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	l := log.NewLogger(cfg.Env)

	defer func() {
		if err := recover(); err != nil {
			l.Error("recovery ❌", zap.Any("error", err))
		}
	}()

	db := scylla.MustConnect(cfg.DB)
	defer db.Close()

	l.Info("db connected ✔")

	repo := scylla.NewRepository(db)

	uc := usecase.NewUsecase(repo)

	srvr := grpc.NewAdapter(uc, l)

	l.Info("layers are set up ✔")

	shutdown(func() {
		l.Info("server is running ✔", zap.Int("port", cfg.Port))
		run(cfg.Port, srvr)
	})

	l.Info("server stopped ✔")
}

func run(port int, s *gRPC.Server) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic("failed to listen tcp: " + err.Error())
	}

	if err = s.Serve(l); err != nil {
		panic("failed to start server: " + err.Error())
	}
}

func shutdown(fn func()) {
	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)
	go fn()
	<-chStop
}
