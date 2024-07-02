package mux

import (
	"encoding/json"
	"net/http"

	"github.com/nautilusgames/demo/wallet/internal/ent"
	"go.uber.org/zap"
)

func New(logger *zap.Logger, entClient *ent.Client) *http.ServeMux {
	// Flag gets printed as a page
	mux := http.NewServeMux()
	// Health endpoint
	mux.HandleFunc("/status", httpHealth())
	mux.HandleFunc("/api/v1/wallet/create-wallet", httpCreateWallet(logger, entClient))
	mux.HandleFunc("/api/v1/wallet/transfer", httpTransfer(logger, entClient))
	mux.HandleFunc("/api/v1/wallet/get-wallet", httpGetWallet(logger, entClient))

	return mux
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
