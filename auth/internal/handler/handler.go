package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/token"
	"github.com/nautilusgames/demo/config/pb"
)

var _expireTokenDuration = 24 * time.Hour

type httpServer struct {
	logger      *zap.Logger
	cfg         *pb.Config
	entClient   *ent.Client
	webToken    token.Maker
	tenantToken token.Maker
}

func New(
	logger *zap.Logger,
	cfg *pb.Config,
	entClient *ent.Client,
	webToken token.Maker,
	tenantToken token.Maker,
) http.Handler {
	mux := http.NewServeMux()

	s := &httpServer{
		logger:      logger,
		cfg:         cfg,
		entClient:   entClient,
		webToken:    webToken,
		tenantToken: tenantToken,
	}

	handler := corsMiddleware(mux)

	mux.HandleFunc("/status", httpHealth())
	mux.HandleFunc("/api/v1/signin", s.handleSignIn())
	mux.HandleFunc("/api/v1/signup", s.handleSignUp())
	mux.HandleFunc("/api/v1/player/verify", s.handleVerifyPlayer())
	mux.HandleFunc("/api/v1/create-tenant-token", s.handleCreateTenantToken())

	return handler
}

func httpHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func readRequest(logger *zap.Logger, r *http.Request, request interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("decode request failed",
			zap.Any("request", &request),
			zap.Error(err))
		return err
	}
	defer r.Body.Close()

	return nil
}

func respond(logger *zap.Logger, w http.ResponseWriter, response interface{}) {
	bytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("marshal response body failed",
			zap.Any("response", response),
			zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bytes); err != nil {
		logger.Error("write message failed", zap.Error(err))
	}
}
