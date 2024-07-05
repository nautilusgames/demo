package mux

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/wallet/internal/ent"
	"github.com/nautilusgames/demo/wallet/model"
)

func httpRollback(logger *zap.Logger, entClient *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("bet")

		var request model.RollbackRequest
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
			request.GameID,
			request.PlayerID,
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
