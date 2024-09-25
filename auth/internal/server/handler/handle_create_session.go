package handler

import (
	"net/http"

	sgbuilder "github.com/nautilusgames/sdk-go/builder"
)

func (h *Handler) HandleCreateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := h.authorize(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var request *CreateSessionRequest
		err = sgbuilder.ToRequest(r.Body, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, _, err := h.tenantPlayerToken.CreateToken(
			request.GameId,
			info.PlayerID,
			info.Username,
			_expireTokenDuration,
		)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}

		sgbuilder.SendReply(w, &CreateSessionResponse{
			Token: token,
		})
	}
}
