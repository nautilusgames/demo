package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/token"
	"github.com/nautilusgames/demo/config"
	pb "github.com/nautilusgames/demo/config/pb"
	"github.com/nautilusgames/demo/wallet/internal/ent"
	"github.com/nautilusgames/demo/wallet/internal/mux"
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

	// connect to mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		cfg.GetDatabase().GetUsername(),
		cfg.GetDatabase().GetPassword(),
		cfg.GetDatabase().GetHost(),
		cfg.GetDatabase().GetPort(),
		cfg.GetDatabase().GetName(),
	)
	logger.Info("---", zap.Any("db", cfg.GetDatabase()), zap.Any("dsn", dsn))
	entClient, err := ent.Open("mysql", dsn)
	if err != nil {
		logger.Fatal("failed opening connection to mysql", zap.Error(err))
	}
	defer entClient.Close()
	if err := entClient.Schema.Create(context.Background()); err != nil {
		logger.Fatal("failed creating schema resources", zap.Error(err))
	}

	tenantPlayerSigning := cfg.GetAuth().GetTenantPlayerSigning()
	tenantPlayerToken, err := token.New(
		tenantPlayerSigning.GetSigningKey(),
		tenantPlayerSigning.GetIssuer(),
		tenantPlayerSigning.GetAudience())
	if err != nil {
		logger.Fatal("failed to create player token", zap.Error(err))
	}

	mux := mux.New(logger, entClient, tenantPlayerToken)
	address := fmt.Sprintf("%s:%d", cfg.Listener.GetTcp().Address, cfg.Listener.GetTcp().Port)
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
