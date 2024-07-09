package httpserver

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/server/httpserver/handler"
	"github.com/nautilusgames/demo/auth/internal/server/httpserver/middleware"
	"github.com/nautilusgames/demo/auth/token"
	"github.com/nautilusgames/demo/config/pb"
)

const (
	_statusPath        = "/status"
	_signInPath        = "/api/v1/signin"
	_signUpPath        = "/api/v1/signup"
	_createSessionPath = "/api/v1/create-session"
)

type HttpServer interface {
	Serve(addr string) error
	Shutdown(ctx context.Context) error
}

type httpServer struct {
	http.Server

	logger *zap.Logger
	mux    *mux.Router
}

func New(
	logger *zap.Logger,
	cfg *pb.Config,
	entClient *ent.Client,
	accessToken token.Maker,
	playerTenantToken token.Maker,
	address string,
) HttpServer {
	mux := mux.NewRouter()
	handler := handler.New(logger, cfg, entClient, accessToken, playerTenantToken)

	// set up routes
	mux.HandleFunc(_statusPath, handler.HandleStatus()).Methods(http.MethodGet)
	mux.HandleFunc(_signInPath, handler.HandleSignIn()).Methods(http.MethodPost)
	mux.HandleFunc(_signUpPath, handler.HandleSignUp()).Methods(http.MethodPost)
	mux.HandleFunc(_createSessionPath, handler.HandleCreateSession()).Methods(http.MethodPost)

	// set up middleware
	mux.Use(middleware.CorsMiddleware)

	return &httpServer{
		logger: logger,
		Server: http.Server{
			Addr:    address,
			Handler: mux,
		},
	}
}

func (h *httpServer) Serve(addr string) error {
	return h.Server.ListenAndServe()
}

func (h *httpServer) Shutdown(ctx context.Context) error {
	return h.Server.Shutdown(ctx)
}
