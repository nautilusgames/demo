package handler

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/checker"
	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/ent/player"
	"github.com/nautilusgames/demo/auth/model"
)

func (s *httpServer) handleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request *model.SignInRequest
		err := readRequest(s.logger, r, &request)
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

		player, err := s.entClient.Player.
			Query().
			Where(player.Username(username)).
			Only(r.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				http.Error(w, "invalid username or password", http.StatusUnauthorized)
				return
			}
			s.logger.Error("failed to query player", zap.Error(err))
			http.Error(w, "failed to query player", http.StatusInternalServerError)
			return
		}

		if err := checker.CheckPassword(password, player.HashedPassword); err != nil {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			return
		}

		token, _, err := s.webToken.CreateToken("", player.ID, player.Username, _expireTokenDuration)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}

		respond(s.logger, w, &model.SignInResponse{
			DisplayName: player.DisplayName,
			Username:    player.Username,
			Currency:    player.Currency,
			Token:       token,
		})
	}
}
