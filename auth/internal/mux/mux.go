package mux

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nautilusgames/demo/auth/internal/checker"
	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/ent/player"
	"github.com/nautilusgames/demo/auth/internal/token"
	"go.uber.org/zap"
)

var _expireTokenDuration = 24 * time.Hour

func New(logger *zap.Logger, entClient *ent.Client, tokenMaker token.Maker) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/status", httpHealth())
	mux.HandleFunc("/api/v1/player/verify", httpAuth(logger))
	mux.HandleFunc("/api/v1/signin", handleSignIn(logger, entClient, tokenMaker))
	mux.HandleFunc("/api/v1/signup", handleSignUp(logger, entClient, tokenMaker))

	return mux
}

func httpAuth(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"status\": \"ok\"}")
	}
}

func httpHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func handleSignIn(logger *zap.Logger, entClient *ent.Client, tokenMaker token.Maker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		if username == "" || password == "" {
			http.Error(w, "username and password are required", http.StatusBadRequest)
			return
		}

		player, err := entClient.Player.
			Query().
			Where(player.Username(username)).
			Only(r.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				http.Error(w, "invalid username or password", http.StatusUnauthorized)
				return
			}
			logger.Error("failed to query player", zap.Error(err))
			http.Error(w, "failed to query player", http.StatusInternalServerError)
			return
		}

		if err := checker.CheckPassword(password, player.HashedPassword); err != nil {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			return
		}

		token, _, err := tokenMaker.CreateToken(player.ID, player.Username, _expireTokenDuration)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"display_name\": \"%s\", \"username\": \"%s\",\"token\": \"%s\"}", player.DisplayName, player.Username, token)
	}
}

func handleSignUp(logger *zap.Logger, entClient *ent.Client, tokenMaker token.Maker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		if username == "" || password == "" {
			http.Error(w, "username and password are required", http.StatusBadRequest)
			return
		}

		hashedPassword, err := checker.HashPassword(password)
		if err != nil {
			http.Error(w, "failed to hash password", http.StatusInternalServerError)
			return
		}

		player, err := entClient.Player.
			Create().
			SetUsername(username).
			SetHashedPassword(hashedPassword).
			SetDisplayName(username).
			Save(r.Context())
		if err != nil {
			logger.Error("failed to create player", zap.Error(err))
			http.Error(w, "failed to create player", http.StatusInternalServerError)
			return
		}

		token, _, err := tokenMaker.CreateToken(player.ID, player.Username, _expireTokenDuration)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"display_name\": \"%s\", \"username\": \"%s\",\"token\": \"%s\"}", player.DisplayName, player.Username, token)
	}
}
