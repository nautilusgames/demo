package httpserver

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	webhook "github.com/nautilusgames/sdk-go/webhook"
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
	_createTenantToken = "/api/v1/create-tenant-token"
)

type HttpServer interface {
	Serve() error
	Shutdown(ctx context.Context) error
}

type httpServer struct {
	http.Server
	logger *zap.Logger
}

func New(
	logger *zap.Logger,
	cfg *pb.Config,
	entClient *ent.Client,
	accessToken token.Maker,
	tenantPlayerToken token.Maker,
	address string,
) HttpServer {
	mux := mux.NewRouter()
	handler := handler.New(logger, cfg, entClient, accessToken, tenantPlayerToken)

	// set up routes
	mux.HandleFunc(_statusPath, handler.HandleStatus()).Methods(http.MethodGet)
	mux.HandleFunc(_signInPath, handler.HandleSignIn()).Methods(http.MethodPost)
	mux.HandleFunc(_signUpPath, handler.HandleSignUp()).Methods(http.MethodPost)
	mux.HandleFunc(_createSessionPath, handler.HandleCreateSession()).Methods(http.MethodPost)
	// deprecated
	mux.HandleFunc(_createTenantToken, handler.HandleCreateTenantToken()).Methods(http.MethodPost)

	// set up webhook
	webhook.HandleVerifyPlayer(mux, logger, handler.HandleVerifyPlayer)

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

func (h *httpServer) Serve() error {
	return h.Server.ListenAndServe()
}

func (h *httpServer) Shutdown(ctx context.Context) error {
	return h.Server.Shutdown(ctx)
}
