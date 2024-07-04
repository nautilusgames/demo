package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/checker"
	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/tx"
	"github.com/nautilusgames/demo/auth/model"
	walletmodel "github.com/nautilusgames/demo/wallet/model"
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
		token := ""

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

		var player *ent.Player
		err = tx.WithTx(r.Context(), s.entClient, func(tx *ent.Tx) error {
			player, err = tx.Player.
				Create().
				SetUsername(username).
				SetHashedPassword(hashedPassword).
				SetDisplayName(displayName).
				SetCurrency(currency).
				Save(r.Context())
			if err != nil {
				s.logger.Error("failed to create player", zap.Error(err))
				return err
			}

			token, _, err = s.accessToken.CreateToken("", player.ID, player.Username, _expireTokenDuration)
			if err != nil {
				s.logger.Error("failed to create token", zap.Error(err))
				return err
			}

			var body bytes.Buffer
			err = json.NewEncoder(&body).Encode(walletmodel.CreateWalletRequest{
				PlayerID: player.ID,
				Currency: currency,
			})
			if err != nil {
				s.logger.Error("failed to encode body", zap.Error(err))
				return err
			}
			url := fmt.Sprintf("%s%s", walletmodel.InternalAddress, walletmodel.CreateWalletPath)
			resp, err := http.Post(url, "application/json", &body)
			if err != nil {
				s.logger.Error("failed to post http request", zap.Error(err))
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				s.logger.Error("failed to create wallet", zap.Any("status", resp.StatusCode))
				return errors.New("failed to create wallet")
			}
			return nil
		})
		if err != nil {
			http.Error(w, "failed to sign up", http.StatusInternalServerError)
		}

		respond(s.logger, w, &model.SignUpResponse{
			DisplayName: player.DisplayName,
			Username:    player.Username,
			Token:       token,
		})
	}
}
