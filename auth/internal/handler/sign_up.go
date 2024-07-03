package handler

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/checker"
)

func (s *httpServer) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		currency := r.PostFormValue("currency")

		if username == "" || password == "" {
			http.Error(w, "username and password are required", http.StatusBadRequest)
			return
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
			SetDisplayName(username).
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"display_name\": \"%s\", \"username\": \"%s\",\"token\": \"%s\"}", player.DisplayName, player.Username, token)
	}
}
