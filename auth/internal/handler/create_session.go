package handler

import (
	"net/http"

	"github.com/nautilusgames/demo/auth/model"
)

func (s *httpServer) handleCreateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := s.authorizeAccessToken(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var request *model.CreatePlayerTenantTokenRequest
		err = readRequest(s.logger, r, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, _, err := s.playerTenantToken.CreateToken(
			request.GameId,
			info.PlayerID,
			info.Username,
			_expireTokenDuration,
		)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}
		respond(s.logger, w, &model.CreatePlayerTenantTokenResponse{
			TenantId: s.cfg.GetTenantId(),
			Token:    token,
		})
	}
}
