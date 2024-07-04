package handler

import (
	"net/http"

	"github.com/nautilusgames/demo/auth/model"
)

func (s *httpServer) handleCreateTenantToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := s.authorize(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var request *model.CreateTenantTokenRequest
		err = readRequest(s.logger, r, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, _, err := s.tenantToken.CreateToken(request.GameId, info.PlayerID, info.Username, _expireTokenDuration)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}
		respond(s.logger, w, &model.CreateTenantTokenResponse{
			Token: token,
		})
	}
}
