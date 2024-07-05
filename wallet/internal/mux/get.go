package mux

import (
	"context"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/wallet/internal/ent"
	entwallet "github.com/nautilusgames/demo/wallet/internal/ent/wallet"
	"github.com/nautilusgames/demo/wallet/model"
)

func httpGet(logger *zap.Logger, entClient *ent.Client) http.HandlerFunc {
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
			value := r.Header.Get("x-player-id")
			id, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthenticated"))
				return
			}
			request.PlayerID = id

		}

		playerWallet, err := GetWallet(r.Context(), entClient, request.PlayerID)
		if err != nil {
			logger.Error("get wallet failed", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		respond(logger, w, model.GetWalletResponse{
			Data:  playerWallet,
			Error: nil,
		})
	}
}

func GetWallet(ctx context.Context, entClient *ent.Client, playerID int64) (*model.PlayerWallet, error) {
	player, err := entClient.Wallet.Query().
		Where(entwallet.ID(playerID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &model.PlayerWallet{
		Balance:  player.Balance,
		LastTxID: player.UpdatedAt.Unix(),
	}, nil
}
