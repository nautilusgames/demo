package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/server/httpserver"
	"github.com/nautilusgames/demo/auth/token"
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

	// connect to mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		cfg.GetDatabase().GetUsername(),
		cfg.GetDatabase().GetPassword(),
		cfg.GetDatabase().GetHost(),
		cfg.GetDatabase().GetPort(),
		cfg.GetDatabase().GetName(),
	)
	entClient, err := ent.Open("mysql", dsn)
	if err != nil {
		logger.Fatal("failed opening connection to mysql", zap.Error(err))
	}
	defer entClient.Close()
	if err := entClient.Schema.Create(context.Background(), schema.WithDropColumn(true)); err != nil {
		logger.Fatal("failed creating schema resources", zap.Error(err))
	}

	playerSigning := cfg.GetAuth().GetPlayerSigning()
	playerToken, err := token.New(playerSigning.GetSigningKey(), playerSigning.GetIssuer(), playerSigning.GetAudience())
	if err != nil {
		logger.Fatal("failed to create player token", zap.Error(err))
	}

	playerTenantSigning := cfg.GetAuth().GetPlayerTenantSigning()
	playerTenantToken, err := token.New(
		playerTenantSigning.GetSigningKey(),
		playerTenantSigning.GetIssuer(),
		playerTenantSigning.GetAudience())
	if err != nil {
		logger.Fatal("failed to create player token", zap.Error(err))
	}

	address := fmt.Sprintf("%s:%d", cfg.Listener.GetTcp().Address, cfg.Listener.GetTcp().Port)
	httpServer := httpserver.New(logger, cfg, entClient, playerToken, playerTenantToken, address)

	serverCh := make(chan struct{})
	go func() {
		logger.Info("server is listening ", zap.String("address", address))
		if err := httpServer.Serve(); err != http.ErrServerClosed {
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

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Fatal("failed to shutdown server", zap.Error(err))
	}

	os.Exit(2)
}
