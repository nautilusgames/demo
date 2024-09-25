package handler

import (
	"context"

	"github.com/nautilusgames/demo/auth/verifier"
	"github.com/nautilusgames/sdk-go/webhook"
	"go.uber.org/zap"
)

func (h *Handler) HandleVerifyPlayer(ctx context.Context, request *webhook.VerifyPlayerRequest) (*webhook.VerifyPlayerReply, error) {
	reply := &webhook.VerifyPlayerReply{}

	payload, webhookErr := verifier.Verify(h.cfg, h.tenantPlayerToken, request.Header)
	if webhookErr != nil {
		h.logger.Error("failed to verify player", zap.Any("error", (webhookErr)))
		reply.Error = webhookErr
		return reply, nil
	}

	player, err := h.entClient.Player.Get(ctx, payload.PlayerID)
	if err != nil {
		h.logger.Error("failed to get player", zap.Error(err))
		return nil, err
	}

	if player.TenantID != request.Header.XTenantId {
		h.logger.Error("player and tenant id mismatch", zap.Error(err))
		return nil, err
	}

	reply.Data = &webhook.PlayerInfo{
		Id:       player.Username,
		Nickname: player.DisplayName,
		Avatar:   "https://s.n11s.io/media/avatar/Ava_F_2_EHD5g5dlsf_2023080210.png",
	}
	return reply, nil
}
