package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/mux"
	"github.com/nautilusgames/demo/config"
	pb "github.com/nautilusgames/demo/config/pb"
)

func Run(f *config.Flags) {
	tmpLogger := config.NewTmpLogger().With(zap.String("filename", f.ConfigPath))

	var cfg pb.Config
	if err := config.ParseFile(f.ConfigPath, &cfg, f.Template); err != nil {
		tmpLogger.Fatal("parsing configuration failed", zap.Error(err))
	}

	RunWithConfig(&cfg)
}

func RunWithConfig(cfg *pb.Config) {
	logger, err := config.NewLogger(cfg.Logger)
	if err != nil {
		config.NewTmpLogger().Fatal("could not instantiate logger", zap.Error(err))
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			panic(err)
		}
	}()

	address := fmt.Sprintf("%s:%d", cfg.Listener.GetTcp().Address, cfg.Listener.GetTcp().Port)
	mux := mux.New(logger)
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	serverCh := make(chan struct{})
	go func() {
		logger.Info("server is listening ", zap.String("address", address))
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal("server exited with", zap.Error(err))
		}
		close(serverCh)
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	<-signalCh

	logger.Info("received interrupt, shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("failed to shutdown server", zap.Error(err))
	}

	os.Exit(2)
}
