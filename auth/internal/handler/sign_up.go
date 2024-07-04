package handler

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/checker"
	"github.com/nautilusgames/demo/auth/internal/model"
)

const _defaultCurrency = "vnd"

func (s *httpServer) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request *model.SignUpRequest
		err := readRequest(s.logger, r, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		username := request.Username
		password := request.Password
		displayName := request.DisplayName
		currency := request.Currency

		if username == "" || password == "" {
			http.Error(w, "username and password are required", http.StatusBadRequest)
			return
		}
		if displayName == "" {
			displayName = username
		}
		if currency == "" {
			currency = _defaultCurrency
		}

		hashedPassword, err := checker.HashPassword(password)
		if err != nil {
			http.Error(w, "failed to hash password", http.StatusInternalServerError)
			return
		}

		player, err := s.entClient.Player.
			Create().
			SetUsername(username).
			SetHashedPassword(hashedPassword).
			SetDisplayName(displayName).
			SetCurrency(currency).
			Save(r.Context())
		if err != nil {
			s.logger.Error("failed to create player", zap.Error(err))
			http.Error(w, "failed to create player", http.StatusInternalServerError)
			return
		}

		token, _, err := s.tokenMaker.CreateToken(player.ID, player.Username, _expireTokenDuration)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}

		respond(s.logger, w, &model.SignUpResponse{
			DisplayName: player.DisplayName,
			Username:    player.Username,
			Token:       token,
		})
	}
}
