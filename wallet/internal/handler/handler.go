package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nautilusgames/sdk-go/webhook"
	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/token"
	"github.com/nautilusgames/demo/wallet/internal/ent"
	entsession "github.com/nautilusgames/demo/wallet/internal/ent/session"
	entwallet "github.com/nautilusgames/demo/wallet/internal/ent/wallet"
	"github.com/nautilusgames/demo/wallet/internal/tx"
)

const (
	_apiKeyHeader      = "x-api-key"
	_tenantIdHeader    = "x-tenant-id"
	_gameIdHeader      = "x-game-id"
	_tenantTokenHeader = "x-tenant-token"
)

var (
	_initWallet int64 = 100000000

	_insufficientBalanceCode  int64 = 1
	_insufficientBalanceError error = errors.New("insufficient balance")
)

func New(logger *zap.Logger, entClient *ent.Client, tokenMaker token.Maker) *Handler {
	return &Handler{
		logger:     logger,
		entClient:  entClient,
		tokenMaker: tokenMaker,
	}
}

type Handler struct {
	logger     *zap.Logger
	entClient  *ent.Client
	tokenMaker token.Maker
}

func (h *Handler) authorizePlayerTenantToken(header *webhook.HookRequestHeader) (*token.Payload, error) {
	// validate headers
	if header.XApiKey == "" {
		return nil, fmt.Errorf("missing api key header")
	}
	if header.XTenantId == "" {
		return nil, fmt.Errorf("missing tenant id header")
	}
	if header.XGameId == "" {
		return nil, fmt.Errorf("missing game id header")
	}
	if header.XTenantToken == "" {
		return nil, fmt.Errorf("missing tenant token header")
	}

	// validate tenant token
	payload, err := h.tokenMaker.VerifyToken(header.XTenantToken)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant token: %s", err)
	}

	if payload.GameID != header.XGameId {
		return nil, fmt.Errorf("invalid game id")
	}

	return payload, nil
}

func (h *Handler) transfer(
	ctx context.Context,
	sessionID int64,
	gameID string,
	playerID int64,
	amount int64,
) (*webhook.WalletTransaction, error) {
	var (
		now         = time.Now()
		transaction *webhook.WalletTransaction
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

		if p.Balance+amount < 0 {
			return _insufficientBalanceError
		} else {
			p.Balance += amount
		}

		walletSessionID, err := getOrCreateSession(ctx, tx.Session, gameID, sessionID)
		if err != nil {
			h.logger.Error("get or create session failed", zap.Error(err))
			return err
		}

		transaction = &webhook.WalletTransaction{
			Id:         now.UnixNano(),
			SessionId:  walletSessionID,
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

func Error(code int64, msg string) *webhook.Error {
	return &webhook.Error{
		Code:    code,
		Message: msg,
	}
}
