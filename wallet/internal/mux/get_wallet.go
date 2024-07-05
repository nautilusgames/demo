package mux

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/wallet/internal/ent"
	entwallet "github.com/nautilusgames/demo/wallet/internal/ent/wallet"
	"github.com/nautilusgames/demo/wallet/model"
)

func httpGetWallet(logger *zap.Logger, entClient *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("get wallet")

		var request model.GetWalletRequest
		if err := readRequest(logger, r, &request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if request.PlayerID == 0 {
			// TODO: get header name from constant
			value := r.Header.Get("X-Player-Id")
			id, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthenticated"))
				return
			}
			request.PlayerID = id

		}

		player, err := entClient.Wallet.Query().
			Where(entwallet.ID(request.PlayerID)).
			Only(r.Context())
		if err != nil {
			logger.Error("get player failed", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		respond(logger, w, model.GetWalletResponse{
			Data: &model.PlayerWallet{
				Balance:  player.Balance,
				LastTxID: player.UpdatedAt.Unix(),
			},
		})
	}
}
