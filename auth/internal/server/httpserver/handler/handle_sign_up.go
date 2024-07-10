package handler

import (
	"net/http"

	sgbuilder "github.com/nautilusgames/sdk-go/builder"
	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/checker"
	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/internal/tx"
	walletcli "github.com/nautilusgames/demo/wallet/client"
)

const _defaultCurrency = "vnd"

func (h *Handler) HandleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request *SignUpRequest
		err := sgbuilder.ToRequest(r.Body, &request)
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
		err = tx.WithTx(r.Context(), h.entClient, func(tx *ent.Tx) error {
			player, err = tx.Player.
				Create().
				SetUsername(username).
				SetHashedPassword(hashedPassword).
				SetDisplayName(displayName).
				SetCurrency(currency).
				Save(r.Context())
			if err != nil {
				h.logger.Error("failed to create player", zap.Error(err))
				return err
			}

			token, _, err = h.accessToken.CreateToken("", player.ID, player.Username, _expireTokenDuration)
			if err != nil {
				h.logger.Error("failed to create token", zap.Error(err))
				return err
			}

			createWalletReq := &walletcli.CreateWalletRequest{
				PlayerID: player.ID,
				Currency: currency,
			}
			_, err := walletcli.CreateWallet(h.logger, createWalletReq)
			if err != nil {
				h.logger.Error("failed to create wallet", zap.Error(err))
				return err
			}

			return nil
		})
		if err != nil {
			http.Error(w, "failed to sign up", http.StatusInternalServerError)
			return
		}

		sgbuilder.SendResponse(w, &SignUpResponse{
			DisplayName: player.DisplayName,
			Username:    player.Username,
			Token:       token,
		})
	}
}
