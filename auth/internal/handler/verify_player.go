package handler

import (
	"fmt"
	"net/http"

	"github.com/nautilusgames/demo/auth/tenant"
	"go.uber.org/zap"
)

type Data struct {
	Data VerifyPlayerResponse `json:"data"`
}
type VerifyPlayerResponse struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (s *httpServer) handleVerifyPlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, playerID, _, err := tenant.GetTenantAuthorization(s.logger, s.cfg, s.playerTenantToken)(w, r)
		if err != nil {
			s.logger.Error("failed to verify tenant token", zap.Error(err))
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		player, err := s.entClient.Player.Get(r.Context(), playerID)
		if err != nil {
			s.logger.Error("failed to get player", zap.Error(err))
			http.Error(w, "failed to get player", http.StatusInternalServerError)
			return
		}

		respond(s.logger, w, &Data{
			Data: VerifyPlayerResponse{
				Id:       fmt.Sprintf("%d", player.ID),
				Nickname: player.DisplayName,
			},
		})
	}
}
