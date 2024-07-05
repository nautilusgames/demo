package mux

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/tenant"
	"github.com/nautilusgames/demo/wallet/internal/ent"
	entwallet "github.com/nautilusgames/demo/wallet/internal/ent/wallet"
	"github.com/nautilusgames/demo/wallet/model"
)

func httpGet(logger *zap.Logger, entClient *ent.Client, tenantAuth tenant.TenantAuthorization) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("get wallet")
		_, playerID, _, err := tenantAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var request model.GetWalletRequest
		if err := readRequest(logger, r, &request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		playerWallet, err := GetWallet(r.Context(), entClient, playerID)
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
