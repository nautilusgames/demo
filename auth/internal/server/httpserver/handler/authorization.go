package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nautilusgames/demo/auth/token"
)

const (
	_authorizationHeader = "authorization"
	_authorizationBearer = "bearer"
)

func (h *Handler) authorizeAccessToken(r *http.Request) (*token.Payload, error) {
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
	payload, err := h.accessToken.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}
