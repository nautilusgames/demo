package handler

import (
	"context"
	"fmt"

	"github.com/nautilusgames/sdk-go/webhook"
	"go.uber.org/zap"
)

func (h *Handler) HandleVerifyPlayer(ctx context.Context, request *webhook.VerifyPlayerRequest) (*webhook.VerifyPlayerResponse, error) {
	var (
		tenantPlayerToken = request.Header.XTenantPlayerToken
		gameID            = request.Header.XGameId
	)
	if tenantPlayerToken == "" {
		return nil, fmt.Errorf("missing tenant player token header")
	}
	if gameID == "" {
		return nil, fmt.Errorf("missing game id header")
	}

	// validate tenant player token
	payload, err := h.tenantPlayerToken.VerifyToken(tenantPlayerToken)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant token: %s", err)
	}

	if gameID != payload.GameID {
		return nil, fmt.Errorf("invalid game id")
	}

	player, err := h.entClient.Player.Get(ctx, payload.PlayerID)
	if err != nil {
		h.logger.Error("failed to get player", zap.Error(err))
		return nil, err
	}

	return &webhook.VerifyPlayerResponse{
		Data: &webhook.PlayerInfo{
			Id:       player.Username,
			Nickname: player.DisplayName,
			Avatar:   "https://s.n11s.io/media/avatar/Ava_F_2_EHD5g5dlsf_2023080210.png",
		},
	}, nil
}
