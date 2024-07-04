package handler

import (
	"net/http"

	"go.uber.org/zap"
)

type VerifyPlayerResponse struct {
	PlayerId    int64  `json:"player_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

func (s *httpServer) handleVerifyPlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := s.authorize(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		player, err := s.entClient.Player.Get(r.Context(), payload.PlayerId)
		if err != nil {
			s.logger.Error("failed to get player", zap.Error(err))
			http.Error(w, "failed to get player", http.StatusInternalServerError)
			return
		}

		respond(s.logger, w, &VerifyPlayerResponse{
			PlayerId:    player.ID,
			Username:    player.Username,
			DisplayName: player.DisplayName,
		})
	}
}
