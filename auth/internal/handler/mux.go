package handler

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/token"
	"github.com/nautilusgames/demo/config/pb"
)

var _expireTokenDuration = 24 * time.Hour

type httpServer struct {
	logger     *zap.Logger
	cfg        *pb.Config
	entClient  *ent.Client
	tokenMaker token.Maker
}

func New(
	logger *zap.Logger,
	cfg *pb.Config,
	entClient *ent.Client,
	tokenMaker token.Maker,
) http.Handler {
	mux := http.NewServeMux()

	s := &httpServer{
		logger:     logger,
		cfg:        cfg,
		entClient:  entClient,
		tokenMaker: tokenMaker,
	}

	handler := corsMiddleware(mux)

	mux.HandleFunc("/status", httpHealth())
	mux.HandleFunc("/api/v1/player/verify", httpAuth(logger))
	mux.HandleFunc("/api/v1/signin", s.handleSignIn())
	mux.HandleFunc("/api/v1/signup", s.handleSignUp())
	mux.HandleFunc("/api/v1/create-tenant-token", s.handleCreateTenantToken())

	return handler
}

func httpAuth(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"status\": \"ok\"}")
	}
}

func httpHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
