package handler

import (
	"fmt"
	"net/http"

	"github.com/nautilusgames/demo/auth/model"
	"github.com/nautilusgames/demo/auth/tenant"
	sgbuilder "github.com/nautilusgames/sdk-go/builder"
	"go.uber.org/zap"
)

func (h *Handler) HandleVerifyPlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, playerID, _, err := tenant.GetTenantAuthorization(h.logger, h.cfg, h.playerTenantToken)(w, r)
		if err != nil {
			h.logger.Error("failed to verify tenant token", zap.Error(err))
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		player, err := h.entClient.Player.Get(r.Context(), playerID)
		if err != nil {
			h.logger.Error("failed to get player", zap.Error(err))
			http.Error(w, "failed to get player", http.StatusInternalServerError)
			return
		}

		sgbuilder.SendResponse(w, &model.VerifyPlayerResponse{
			Data: &model.PlayerInfo{
				Id:       fmt.Sprintf("%d", player.ID),
				Nickname: player.DisplayName,
			},
		})
	}
}
