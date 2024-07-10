package mux

import (
	"fmt"
	"net/http"

	"github.com/nautilusgames/demo/auth/token"
)

const (
	_apiKeyHeader      = "x-api-key"
	_tenantIdHeader    = "x-tenant-id"
	_gameIdHeader      = "x-game-id"
	_tenantTokenHeader = "x-tenant-token"
)

func authorizePlayerTenantToken(r *http.Request, tokenMaker token.Maker) (*token.Payload, error) {
	// validate headers
	apiKey := r.Header.Get(_apiKeyHeader)
	if apiKey == "" {
		return nil, fmt.Errorf("missing api key header")
	}
	tenantId := r.Header.Get(_tenantIdHeader)
	if tenantId == "" {
		return nil, fmt.Errorf("missing tenant id header")
	}
	gameId := r.Header.Get(_gameIdHeader)
	if gameId == "" {
		return nil, fmt.Errorf("missing game id header")
	}
	tenantToken := r.Header.Get(_tenantTokenHeader)
	if tenantToken == "" {
		return nil, fmt.Errorf("missing tenant token header")
	}

	// validate tenant token
	payload, err := tokenMaker.VerifyToken(tenantToken)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant token: %s", err)
	}

	if payload.GameID != gameId {
		return nil, fmt.Errorf("invalid game id")
	}

	return payload, nil
}
