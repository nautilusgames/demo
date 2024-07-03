package mux

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/wallet/internal/ent"
	"github.com/nautilusgames/demo/wallet/model"
)

func httpCreateWallet(logger *zap.Logger, entClient *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("create wallet")

		var request model.CreateWalletRequest
		if err := readRequest(logger, r, &request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if request.Currency == "" || request.PlayerID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request"))
			return
		}

		err := entClient.Wallet.Create().
			SetID(request.PlayerID).
			SetCurrency(request.Currency).
			Exec(r.Context())
		if err != nil {
			logger.Error("create new player failed", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		respond(logger, w, model.CreateWalletResponse{})
	}
}
