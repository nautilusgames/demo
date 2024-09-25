package handler

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/verifier"
	"github.com/nautilusgames/sdk-go/webhook"
)

func (h *Handler) HandlePayout(ctx context.Context, request *webhook.PayoutRequest) (*webhook.TransactionReply, error) {
	reply := &webhook.TransactionReply{}
	payload, webhookErr := verifier.Verify(h.cfg, h.token, request.Header)
	if webhookErr != nil {
		reply.Error = webhookErr
		return reply, nil
	}

	if request.SessionId == "" {
		reply.Error = Error(webhook.ErrInvalidRequest, "invalid session_id")
		return reply, nil
	}

	if request.Amount < 0 {
		reply.Error = Error(webhook.ErrInvalidRequest, "invalid amount")
		return reply, nil
	}

	if request.Amount == 0 {
		playerWallet, err := h.getWallet(ctx, payload.PlayerID)
		if err != nil {
			h.logger.Error("get wallet failed", zap.Error(err))
			reply.Error = Error(webhook.ErrInternalServerError, err.Error())
			return reply, nil
		}

		reply.Data = &webhook.TransactionData{
			TenantTxId:      uuid.NewString(),
			TenantSessionId: request.SessionId,
			Currency:        playerWallet.Currency,
			Amount:          request.Amount,
			NewBalance:      playerWallet.Balance,
			CreatedAt:       playerWallet.UpdatedAt,
		}

		return reply, nil
	}

	tx, err := h.transfer(ctx, request.SessionId, payload.Object, payload.PlayerID, request.Amount)
	if err != nil {
		reply.Error = Error(webhook.ErrInternalServerError, err.Error())
		return reply, nil
	}

	reply.Data = tx
	return reply, nil
}
