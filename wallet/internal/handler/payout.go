package handler

import (
	"context"
	"net/http"

	"github.com/nautilusgames/sdk-go/webhook"
	"go.uber.org/zap"
)

func (h *Handler) HandlePayout(ctx context.Context, request *webhook.PayoutRequest) (*webhook.WalletResponse, error) {
	response := &webhook.WalletResponse{}
	payload, err := h.authorizePlayerTenantToken(request.Header)
	if err != nil {
		response.Error = Error(http.StatusUnauthorized, err.Error())
		return response, nil
	}

	if request.SessionId <= 0 {
		response.Error = Error(http.StatusBadRequest, "invalid session_id")
		return response, nil
	}

	if request.Amount < 0 {
		response.Error = Error(http.StatusBadRequest, "invalid amount")
		return response, nil
	}

	if request.Amount == 0 {
		playerWallet, err := h.getWallet(ctx, payload.PlayerID)
		if err != nil {
			h.logger.Error("get wallet failed", zap.Error(err))
			response.Error = Error(http.StatusInternalServerError, err.Error())
			return response, nil
		}

		tx := &webhook.WalletTransaction{
			Id:         playerWallet.LastTxId,
			NewBalance: playerWallet.Balance,
			SessionId:  request.SessionId,
			Amount:     request.Amount,
		}

		response.Data = tx
		return response, nil
	}

	tx, err := h.transfer(ctx, request.SessionId, payload.GameID, payload.PlayerID, request.Amount)
	if err != nil {
		response.Error = Error(http.StatusInternalServerError, err.Error())
		return response, nil
	}

	response.Data = tx
	return response, nil
}
