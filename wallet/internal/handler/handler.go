package handler

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/nautilusgames/sdk-go/webhook"
	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/token"
	pb "github.com/nautilusgames/demo/config/pb"
	"github.com/nautilusgames/demo/wallet/internal/ent"
	entsession "github.com/nautilusgames/demo/wallet/internal/ent/session"
	entwallet "github.com/nautilusgames/demo/wallet/internal/ent/wallet"
	"github.com/nautilusgames/demo/wallet/internal/tx"
)

var (
	initWallet      int64 = 10000000000
	errInsufficient error = errors.New("insufficient balance")
)

func New(logger *zap.Logger, cfg *pb.Config, entClient *ent.Client, token token.Maker) *Handler {
	return &Handler{
		cfg:       cfg,
		logger:    logger,
		entClient: entClient,
		token:     token,
	}
}

type Handler struct {
	cfg       *pb.Config
	logger    *zap.Logger
	entClient *ent.Client
	token     token.Maker
}

func Error(code int64, msg string) *webhook.Error {
	return &webhook.Error{
		Code:    code,
		Message: msg,
	}
}

func (h *Handler) transfer(
	ctx context.Context,
	sessionID string,
	gameID string,
	playerID int64,
	amount float64,
) (*webhook.TransactionData, error) {
	var (
		now            = time.Now()
		transferAmount = toInternalAmount(amount)
		transaction    *webhook.TransactionData
	)
	err := tx.WithTx(ctx, h.entClient, func(tx *ent.Tx) error {
		p, err := tx.Wallet.Query().
			Where(entwallet.ID(playerID)).
			ForUpdate().
			Only(ctx)
		if err != nil {
			h.logger.Error("get player failed", zap.Error(err))
			return err
		}

		if p.Balance+transferAmount < 0 {
			return errInsufficient
		} else {
			p.Balance += transferAmount
		}

		walletSessionID, err := getOrCreateSession(ctx, tx.Session, gameID, sessionID)
		if err != nil {
			h.logger.Error("get or create session failed", zap.Error(err))
			return err
		}

		transaction = &webhook.TransactionData{
			TenantTxId:      uuid.NewString(),
			TenantSessionId: fmt.Sprint(walletSessionID),
			Amount:          math.Abs(amount),
			NewBalance:      toExternalAmount(p.Balance),
			CreatedAt:       now.UnixNano(),
		}

		if amount == 0 {
			return nil
		}

		err = tx.Wallet.Update().
			Where(entwallet.ID(p.ID)).
			SetBalance(p.Balance).
			Exec(ctx)
		if err != nil {
			h.logger.Error("update player failed", zap.Error(err))
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
	gameSessionID string,
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
	gameSessionID string,
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

const _denominator = 1000

func toInternalAmount(amount float64) int64 {
	return int64(amount * _denominator)
}
func toExternalAmount(amount int64) float64 {
	return float64(amount) / _denominator
}
