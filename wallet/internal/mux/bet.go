package mux

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/token"
	"github.com/nautilusgames/demo/wallet/internal/ent"
	"github.com/nautilusgames/demo/wallet/model"
)

func httpBet(logger *zap.Logger, entClient *ent.Client, tokenMaker token.Maker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("bet")

		payload, err := authorizePlayerTenantToken(r, tokenMaker)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		gameID := payload.GameID
		playerID := payload.PlayerID

		var request model.BetRequest
		if err := readRequest(logger, r, &request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if request.SessionID <= 0 {
			http.Error(w, "invalid session_id", http.StatusBadRequest)
			return
		}

		tx, err := transfer(
			r.Context(),
			entClient,
			logger,
			request.SessionID,
			gameID,
			playerID,
			-request.Amount,
		)
		if err != nil {
			if errors.Is(err, _insufficientBalanceError) {
				respond(logger, w, model.Response{
					Data: nil,
					Error: &model.Error{
						Code:    int32(_insufficientBalanceCode),
						Message: _insufficientBalanceError.Error(),
					},
				})
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respond(logger, w, model.Response{
			Data:  tx,
			Error: nil,
		})
	}
}
