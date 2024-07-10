package mux

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/token"
	"github.com/nautilusgames/demo/wallet/internal/ent"
	entsession "github.com/nautilusgames/demo/wallet/internal/ent/session"
	entwallet "github.com/nautilusgames/demo/wallet/internal/ent/wallet"
	"github.com/nautilusgames/demo/wallet/internal/tx"
	"github.com/nautilusgames/demo/wallet/model"
)

var (
	_insufficientBalanceCode  = 1
	_insufficientBalanceError = errors.New("insufficient balance")
)

func New(logger *zap.Logger, entClient *ent.Client, tokenMaker token.Maker) *http.ServeMux {
	// Flag gets printed as a page
	mux := http.NewServeMux()
	// Health endpoint
	mux.HandleFunc("/status", httpHealth())
	mux.HandleFunc(model.CreatePath, httpCreate(logger, entClient))
	mux.HandleFunc(model.BetPath, httpBet(logger, entClient, tokenMaker))
	mux.HandleFunc(model.PayoutPath, httpPayout(logger, entClient, tokenMaker))
	mux.HandleFunc(model.RefundPath, httpRefund(logger, entClient, tokenMaker))
	mux.HandleFunc(model.GetPath, httpGet(logger, entClient, tokenMaker))

	return mux
}

func httpHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func readRequest(logger *zap.Logger, r *http.Request, request interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("decode request failed",
			zap.Any("request", &request),
			zap.Error(err))
		return err
	}
	defer r.Body.Close()

	return nil
}

func respond(logger *zap.Logger, w http.ResponseWriter, response interface{}) {
	bytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("marshal response body failed",
			zap.Any("response", response),
			zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bytes); err != nil {
		logger.Error("write message failed", zap.Error(err))
	}
}

func transfer(
	ctx context.Context,
	entClient *ent.Client,
	logger *zap.Logger,
	sessionID int64,
	gameID string,
	playerID int64,
	amount int64,
) (*model.Transaction, error) {
	var (
		now         = time.Now()
		transaction *model.Transaction
	)
	err := tx.WithTx(ctx, entClient, func(tx *ent.Tx) error {
		p, err := tx.Wallet.Query().
			Where(entwallet.ID(playerID)).
			ForUpdate().
			Only(ctx)
		if err != nil {
			logger.Error("get player failed", zap.Error(err))
			return err
		}

		if p.Balance+amount < 0 {
			return _insufficientBalanceError
		} else {
			p.Balance += amount
		}

		walletSessionID, err := getOrCreateSession(ctx, tx.Session, gameID, sessionID)
		if err != nil {
			logger.Error("get or create session failed", zap.Error(err))
			return err
		}

		transaction = &model.Transaction{
			ID:         now.UnixNano(),
			SessionID:  walletSessionID,
			Amount:     amount,
			NewBalance: p.Balance,
			CreatedAt:  now.UnixNano(),
		}

		if amount == 0 {
			return nil
		}

		err = tx.Wallet.Update().
			Where(entwallet.ID(p.ID)).
			SetBalance(p.Balance).
			Exec(ctx)
		if err != nil {
			logger.Error("update player failed", zap.Error(err))
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func getOrCreateSession(
	ctx context.Context,
	entSession *ent.SessionClient,
	gameID string,
	gameSessionID int64,
) (int64, error) {
	session, err := entSession.Query().
		Where(
			entsession.GameID(gameID),
			entsession.GameSessionID(gameSessionID),
		).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return 0, errors.New("get session failed")
	}

	if session == nil {
		return createSession(ctx, entSession, gameID, gameSessionID)
	}

	return session.ID, nil
}

func createSession(
	ctx context.Context,
	entSession *ent.SessionClient,
	gameID string,
	gameSessionID int64,
) (int64, error) {
	session, err := entSession.Create().
		SetGameID(gameID).
		SetGameSessionID(gameSessionID).
		Save(ctx)
	if err != nil {
		return 0, errors.New("create session failed")
	}

	return session.ID, nil
}
