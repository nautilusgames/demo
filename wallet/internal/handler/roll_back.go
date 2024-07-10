package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/nautilusgames/sdk-go/webhook"
)

func (h *Handler) HandleRollBack(ctx context.Context, request *webhook.BetRequest) (*webhook.WalletResponse, error) {
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

	tx, err := h.transfer(ctx, request.SessionId, payload.GameID, payload.PlayerID, -request.Amount)
	if err != nil {
		if errors.Is(err, _insufficientBalanceError) {
			response.Error = Error(_insufficientBalanceCode, err.Error())
			return response, nil
		}

		response.Error = Error(http.StatusInternalServerError, err.Error())
		return response, nil
	}

	response.Data = tx
	return response, nil
}
