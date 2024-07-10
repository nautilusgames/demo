package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	sgbuilder "github.com/nautilusgames/sdk-go/builder"
	"go.uber.org/zap"
)

const (
	_apiKeyHeader   = "x-api-key"
	_tenantIdHeader = "x-tenant-id"
)

func (h *Handler) HandleCreateTenantToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := h.authorizeAccessToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		tenantToken, err := h.createToken(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sgbuilder.SendResponse(w, &CreateTenantTokenResponse{
			TenantId: h.cfg.GetTenantId(),
			Token:    tenantToken,
		})
	}
}

func (h *Handler) createToken(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.cfg.GetTenantTokenUrl(), nil)
	if err != nil {
		h.logger.Error("create request failed", zap.Error(err))
		return "", err
	}
	req.Header.Set(_tenantIdHeader, fmt.Sprintf("%d", h.cfg.GetTenantId()))
	req.Header.Set(_apiKeyHeader, h.cfg.GetTenantApiKey())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		h.logger.Error("request failed", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		h.logger.Error("request failed", zap.Int("status_code", resp.StatusCode))
		return "", errors.New(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		h.logger.Error("read response failed", zap.Error(err))
		return "", err
	}

	response := map[string]string{}
	if err = json.Unmarshal(data, &response); err != nil {
		h.logger.Error("unmarshal response failed",
			zap.Any("response", string(data)),
			zap.Error(err))
		return "", err
	}

	if len(response["token"]) == 0 {
		h.logger.Error("no token in response",
			zap.Any("response", string(data)))
		return "", errors.New("empty token")
	}

	return response["token"], nil
}
