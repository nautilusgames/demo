package handler

import (
	"context"

	"github.com/nautilusgames/sdk-go/webhook"
	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/verifier"
	entwallet "github.com/nautilusgames/demo/wallet/internal/ent/wallet"
)

func (h *Handler) HandleGetWallet(ctx context.Context, request *webhook.GetWalletRequest) (*webhook.GetWalletReply, error) {
	reply := &webhook.GetWalletReply{}
	payload, webhookErr := verifier.Verify(h.cfg, h.token, request.Header)
	if webhookErr != nil {
		reply.Error = webhookErr
		return reply, nil
	}

	playerWallet, err := h.getWallet(ctx, payload.PlayerID)
	if err != nil {
		h.logger.Error("get wallet failed", zap.Error(err))
		reply.Error = Error(webhook.ErrInternalServerError, err.Error())
		return reply, nil
	}

	reply.Data = playerWallet
	return reply, nil
}

func (h *Handler) getWallet(ctx context.Context, playerID int64) (*webhook.PlayerWallet, error) {
	player, err := h.entClient.Wallet.Query().
		Where(entwallet.ID(playerID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &webhook.PlayerWallet{
		Currency:  player.Currency,
		Balance:   toExternalAmount(player.Balance),
		UpdatedAt: player.UpdatedAt.Unix(),
	}, nil
}
