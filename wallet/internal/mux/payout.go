package mux

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/token"
	"github.com/nautilusgames/demo/wallet/internal/ent"
	"github.com/nautilusgames/demo/wallet/model"
)

func httpPayout(logger *zap.Logger, entClient *ent.Client, tokenMaker token.Maker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("payout")

		payload, err := authorizePlayerTenantToken(r, tokenMaker)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		playerID := payload.PlayerID
		gameID := payload.GameID

		var request model.PayoutRequest
		if err := readRequest(logger, r, &request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if request.SessionID <= 0 {
			http.Error(w, "invalid session_id", http.StatusBadRequest)
			return
		}

		if request.Amount < 0 {
			http.Error(w, "invalid amount", http.StatusBadRequest)
			return
		}

		if request.Amount == 0 {
			playerWallet, err := GetWallet(r.Context(), entClient, playerID)
			if err != nil {
				logger.Error("get wallet failed", zap.Error(err))
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tx := &model.Transaction{
				ID:         playerWallet.LastTxID,
				NewBalance: playerWallet.Balance,
				SessionID:  request.SessionID,
				Amount:     request.Amount,
			}

			respond(logger, w, model.Response{
				Data:  tx,
				Error: nil,
			})
			return
		}

		tx, err := transfer(
			r.Context(),
			entClient,
			logger,
			request.SessionID,
			gameID,
			playerID,
			request.Amount,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respond(logger, w, model.Response{
			Data:  tx,
			Error: nil,
		})
	}
}
