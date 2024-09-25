package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nautilusgames/demo/auth/token"
)

func (h *Handler) authorize(r *http.Request) (*token.Payload, error) {
	value := r.Header.Get("authorization")
	if value == "" {
		return nil, fmt.Errorf("missing authorization header")
	}

	fields := strings.Fields(value)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != "bearer" {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := h.accessToken.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}
