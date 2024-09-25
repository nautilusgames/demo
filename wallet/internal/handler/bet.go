package handler

import (
	"context"
	"errors"

	"github.com/nautilusgames/demo/auth/verifier"
	"github.com/nautilusgames/sdk-go/webhook"
)

func (h *Handler) HandleBet(ctx context.Context, request *webhook.BetRequest) (*webhook.TransactionReply, error) {
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

	tx, err := h.transfer(ctx, request.SessionId, payload.Object, payload.PlayerID, -request.Amount)
	if err != nil {
		if errors.Is(err, errInsufficient) {
			reply.Error = Error(webhook.ErrInsufficient, err.Error())
			return reply, nil
		}

		reply.Error = Error(webhook.ErrInternalServerError, err.Error())
		return reply, nil
	}

	reply.Data = tx
	return reply, nil
}
