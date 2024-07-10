package handler

import (
	"net/http"
	"time"

	sgbuilder "github.com/nautilusgames/sdk-go/builder"
	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/checker"
	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/ent/player"
)

const _expireTokenDuration = 24 * time.Hour

func (h *Handler) HandleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request *SignInRequest
		err := sgbuilder.ToRequest(r.Body, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		username := request.Username
		password := request.Password
		if username == "" || password == "" {
			http.Error(w, "username and password are required", http.StatusBadRequest)
			return
		}

		player, err := h.entClient.Player.
			Query().
			Where(player.Username(username)).
			Only(r.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				http.Error(w, "invalid username or password", http.StatusUnauthorized)
				return
			}

			h.logger.Error("failed to query player", zap.Error(err))
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		if err := checker.CheckPassword(password, player.HashedPassword); err != nil {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			return
		}

		token, _, err := h.accessToken.CreateToken("", player.ID, player.Username, _expireTokenDuration)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}

		sgbuilder.SendResponse(w, &SignInResponse{
			DisplayName: player.DisplayName,
			Username:    player.Username,
			Currency:    player.Currency,
			Token:       token,
		})
	}
}
