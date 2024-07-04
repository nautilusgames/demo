package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/nautilusgames/demo/auth/internal/token"
)

const (
	_authorizationHeader = "authorization"
	_authorizationBearer = "bearer"

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

func (s *httpServer) authorizeAccessToken(_ http.ResponseWriter, r *http.Request) (*token.Payload, error) {
	value := r.Header.Get(_authorizationHeader)
	if value == "" {
		return nil, fmt.Errorf("missing authorization header")
	}

	fields := strings.Fields(value)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != _authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := s.accessToken.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}

func (s *httpServer) authorizeTenantToken(_ http.ResponseWriter, r *http.Request) (tenantID string, playerID int64, gameID string, err error) {
	headers := &Headers{
		TenantID:     r.Header.Get(HeaderTenantID),
		TenantSecret: r.Header.Get(HeaderTenantSecret),
		TenantToken:  r.Header.Get(HeaderTenantToken),
		GameID:       r.Header.Get(HeaderGameID),
	}

	tenantIDNumber, err := strconv.ParseInt(headers.TenantID, 10, 64)
	if err != nil || tenantIDNumber == 0 {
		return "", 0, "", errors.New("invalid tenant id " + tenantID)
	}

	if tenantIDNumber != s.cfg.GetTenantId() || headers.TenantSecret != s.cfg.GetTenantApiKey() {
		return "", 0, "", errors.New("invalid tenant credentials")
	}

	if len(headers.TenantToken) == 0 {
		return "", 0, "", errors.New("unauthorized")
	}

	payload, err := s.tenantToken.VerifyToken(headers.TenantToken)
	if err != nil {
		return "", 0, "", errors.New("unauthorized")
	}

	if payload.GameID != headers.GameID {
		return "", 0, "", errors.New("unauthorized")
	}

	return headers.TenantID, payload.PlayerID, headers.GameID, nil
}
