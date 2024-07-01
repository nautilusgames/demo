package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/nautilusgames/demo/api/v1/demo"
	"github.com/nautilusgames/demo/internal/logger"
	demoServer "github.com/nautilusgames/demo/internal/server/demo"
	config "github.com/nautilusgames/demo/pkg/config"
)

func Run(f *Flags) {
	cfg := loadConfig(f)

	tracer.Start()
	defer tracer.Stop()

	Serve(cfg)
}

// Serve ...
func Serve(cfg *config.Config) {
	_ = context.Background()
	logger := logger.New()

	// http
	mux := httptrace.NewServeMux()
	httpLis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.HttpListener.GetTcp().Address, cfg.HttpListener.GetTcp().Port))
	//grpc
	grpcServer := grpc.NewServer()
	grpcLis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GrpcListener.GetTcp().Address, cfg.GrpcListener.GetTcp().Port))
	if err != nil {
		panic(err)
	}

	// -----------------------------------------------
	// register services
	demo.RegisterDemoServer(grpcServer, demoServer.NewServer())

	// -------------------------------------------------

	reflection.Register(grpcServer)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		logger.Info("http is listening ", zap.String("address", httpLis.Addr().String()))
		if err := http.Serve(httpLis, mux); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()
	go func() {
		defer wg.Done()
		logger.Info("grpc listening", zap.String("address", grpcLis.Addr().String()))
		if err := grpcServer.Serve(grpcLis); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()
	go watchShutdownSignal(logger)

	wg.Wait()
}

func watchShutdownSignal(logger *zap.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan

	logger.Info(fmt.Sprint("got signal:", sig))
	logger.Info("start process before stop")

	os.Exit(0)
}
