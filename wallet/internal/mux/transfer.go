package mux

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/wallet/internal/ent"
	entwallet "github.com/nautilusgames/demo/wallet/internal/ent/wallet"
	"github.com/nautilusgames/demo/wallet/internal/tx"
	"github.com/nautilusgames/demo/wallet/model"
)

func httpTransfer(logger *zap.Logger, entClient *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("transfer")

		var request model.TransferRequest
		if err := readRequest(logger, r, &request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if request.PlayerID == 0 || request.SessionID == 0 || request.GameID == "" || request.Amount == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request"))
			return
		}

		var (
			now         = time.Now()
			sessionID   = request.SessionID
			gameID      = request.GameID
			transaction model.Transaction
		)
		err := tx.WithTx(r.Context(), entClient, func(tx *ent.Tx) error {
			p, err := tx.Wallet.Query().
				Where(entwallet.ID(request.PlayerID)).
				ForUpdate().
				Only(r.Context())
			if err != nil {
				logger.Error("get player failed", zap.Error(err))
				return err
			}

			amount := request.Amount
			if p.Balance+amount < 0 {
				return errors.New("insufficient balance")
			} else {
				p.Balance += amount
			}

			sessionID, gameID, err = getOrCreateSession(r.Context(), logger, tx.Session, sessionID, gameID)
			if err != nil {
				logger.Error("get or create session failed", zap.Error(err))
				return err
			}

			transaction = model.Transaction{
				TxID:      now.UnixNano(),
				CreatedAt: now.UnixNano(),
				PlayerID:  p.ID,
				GameID:    gameID,
				Amount:    amount,
			}

			if amount == 0 {
				return nil
			}

			err = tx.Wallet.Update().
				Where(entwallet.ID(p.ID)).
				SetBalance(p.Balance).
				Exec(r.Context())
			if err != nil {
				logger.Error("update player failed", zap.Error(err))
				return err
			}

			return nil
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		respond(logger, w, model.TransferResponse{
			SessionID: sessionID,
			Tx:        transaction,
		})
		return
	}
}

func getOrCreateSession(
	ctx context.Context,
	logger *zap.Logger,
	entSession *ent.SessionClient,
	sessionID int64,
	gameID string,
) (int64, string, error) {
	if sessionID == 0 {
		return Create(ctx, logger, entSession, gameID)
	}

	session, err := entSession.Get(ctx, sessionID)
	if err != nil {
		logger.Error("get session failed", zap.Error(err))
		return 0, "", errors.New("get session failed")
	}

	if session.GameID != gameID {
		logger.Error("session game id mismatch")
		return 0, "", errors.New("session service id mismatch")
	}

	return session.ID, session.GameID, nil
}

func Create(
	ctx context.Context,
	logger *zap.Logger,
	entSession *ent.SessionClient,
	GameID string,
) (int64, string, error) {
	session, err := entSession.Create().
		SetGameID(GameID).
		Save(ctx)
	if err != nil {
		logger.Error("create session failed", zap.Error(err))
		return 0, "", errors.New("create session failed")
	}

	return session.ID, session.GameID, nil
}
