package tenant

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/nautilusgames/demo/auth/token"
	"github.com/nautilusgames/demo/config/pb"
	"go.uber.org/zap"
)

const (
	HeaderTenantID     = "x-tenant-id"
	HeaderTenantSecret = "x-tenant-secret"
	HeaderTenantToken  = "x-tenant-token"
	HeaderGameID       = "x-game-id"
)

type Headers struct {
	TenantID     string
	TenantSecret string
	TenantToken  string
	GameID       string
}

type TenantAuthorization func(_ http.ResponseWriter, r *http.Request) (tenantID string, playerID int64, gameID string, err error)

func GetTenantAuthorization(logger *zap.Logger, cfg *pb.Config, tokenMaker token.Maker) TenantAuthorization {
	return func(_ http.ResponseWriter, r *http.Request) (tenantID string, playerID int64, gameID string, err error) {
		headers := &Headers{
			TenantID:     r.Header.Get(HeaderTenantID),
			TenantSecret: r.Header.Get(HeaderTenantSecret),
			TenantToken:  r.Header.Get(HeaderTenantToken),
			GameID:       r.Header.Get(HeaderGameID),
		}
		logger.Info("headers", zap.Any("headers", headers))

		tenantIDNumber, err := strconv.ParseInt(headers.TenantID, 10, 64)
		if err != nil || tenantIDNumber == 0 {
			logger.Error("tenant_id not found", zap.Error(err))
			return "", 0, "", errors.New("invalid tenant id " + tenantID)
		}

		if len(headers.TenantToken) == 0 {
			logger.Error("tenant_token not found")
			return "", 0, "", errors.New("unauthorized ")
		}

		payload, err := tokenMaker.VerifyToken(headers.TenantToken)
		if err != nil {
			logger.Error("fail to verify tenant_token")
			return "", 0, "", errors.New("unauthorized")
		}

		if payload.GameID != headers.GameID {
			logger.Error("game id")
			return "", 0, "", errors.New("unauthorized")
		}

		return headers.TenantID, payload.PlayerID, headers.GameID, nil
	}
}
