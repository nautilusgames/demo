package mux

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/wallet/internal/ent"
	"github.com/nautilusgames/demo/wallet/model"
)

func httpRefund(logger *zap.Logger, entClient *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("bet")

		var request model.RefundRequest
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
			playerWallet, err := GetWallet(r.Context(), entClient, request.PlayerID)
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
			request.GameID,
			request.PlayerID,
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
