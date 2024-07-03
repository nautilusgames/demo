package mux

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/checker"
	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/ent/player"
	"github.com/nautilusgames/demo/auth/internal/token"
	"github.com/nautilusgames/demo/auth/internal/tx"
	walletModel "github.com/nautilusgames/demo/wallet/model"
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
		currency := r.PostFormValue("currency")
		token := ""

		if username == "" || password == "" {
			http.Error(w, "username and password are required", http.StatusBadRequest)
			return
		}

		hashedPassword, err := checker.HashPassword(password)
		if err != nil {
			http.Error(w, "failed to hash password", http.StatusInternalServerError)
			return
		}
		err = tx.WithTx(r.Context(), entClient, func(tx *ent.Tx) error {
			player, err := tx.Player.
				Create().
				SetUsername(username).
				SetHashedPassword(hashedPassword).
				SetCurrency(currency).
				SetDisplayName(username).
				Save(r.Context())
			if err != nil {
				logger.Error("failed to create player", zap.Error(err))
				return err
			}

			token, _, err = tokenMaker.CreateToken(player.ID, player.Username, _expireTokenDuration)
			if err != nil {
				return err
			}

			var body bytes.Buffer
			err = json.NewEncoder(&body).Encode(walletModel.CreateWalletRequest{
				PlayerID: player.ID,
				Currency: currency,
			})
			if err != nil {
				logger.Error("failed to encode body", zap.Error(err))
				return err
			}

			url := fmt.Sprintf("%s%s", walletModel.InternalAddress, walletModel.CreateWalletPath)
			resp, err := http.Post(url, "application/json", &body)
			if err != nil {
				logger.Error("failed to post", zap.Error(err))
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				logger.Error("status code", zap.Any("code", resp.StatusCode))
				return errors.New("failed to create wallet")
			}

			return nil
		})
		if err != nil {
			http.Error(w, "failed to create player", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"display_name\": \"%s\", \"username\": \"%s\",\"token\": \"%s\"}", player.DisplayName, player.Username, token)
	}
}
