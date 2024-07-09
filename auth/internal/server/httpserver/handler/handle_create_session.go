package handler

import (
	"net/http"

	sgbuilder "github.com/nautilusgames/sdk-go/builder"

	"github.com/nautilusgames/demo/auth/model"
)

func (h *Handler) HandleCreateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := h.authorizeAccessToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var request *model.CreateSessionRequest
		err = sgbuilder.ToRequest(r.Body, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, _, err := h.playerTenantToken.CreateToken(
			request.GameId,
			info.PlayerID,
			info.Username,
			_expireTokenDuration,
		)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}

		sgbuilder.SendResponse(w, &model.CreateSessionResponse{
			TenantId: h.cfg.GetTenantId(),
			Token:    token,
		})
	}
}
