package handler

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateWalletRequest
	if err := readRequest(h.logger, r, &request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if request.Currency == "" || request.PlayerID == 0 {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := h.entClient.Wallet.Create().
		SetID(request.PlayerID).
		SetCurrency(request.Currency).
		SetBalance(_initWallet).
		Exec(r.Context())
	if err != nil {
		h.logger.Error("create new player failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respond(h.logger, w, CreateWalletResponse{})
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
