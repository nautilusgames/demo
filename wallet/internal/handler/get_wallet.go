package handler

import (
	"context"
	"net/http"

	"github.com/nautilusgames/sdk-go/webhook"
	"go.uber.org/zap"

	entwallet "github.com/nautilusgames/demo/wallet/internal/ent/wallet"
)

func (h *Handler) HandleGetWallet(ctx context.Context, request *webhook.GetWalletRequest) (*webhook.GetWalletResponse, error) {
	response := &webhook.GetWalletResponse{}
	payload, err := h.authorizePlayerTenantToken(request.Header)
	if err != nil {
		response.Error = Error(http.StatusUnauthorized, err.Error())
		return response, nil
	}

	playerWallet, err := h.getWallet(ctx, payload.PlayerID)
	if err != nil {
		h.logger.Error("get wallet failed", zap.Error(err))
		response.Error = Error(http.StatusInternalServerError, err.Error())
		return response, nil
	}

	response.Data = playerWallet
	return response, nil
}

func (h *Handler) getWallet(ctx context.Context, playerID int64) (*webhook.PlayerWallet, error) {
	player, err := h.entClient.Wallet.Query().
		Where(entwallet.ID(playerID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &webhook.PlayerWallet{
		Balance:  player.Balance,
		LastTxId: player.UpdatedAt.Unix(),
	}, nil
}
