package mux

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func (s *httpServer) handleCreateTenantToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := s.authorize(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		tenantToken, err := s.createToken(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"tenant-token\": \"%s\"}", tenantToken)
	}
}

func (s *httpServer) createToken(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.cfg.GetTenantTokenUrl(), nil)
	if err != nil {
		s.logger.Error("create request failed", zap.Error(err))
		return "", err
	}
	req.Header.Set("x-tenant-id", s.cfg.GetTenantId())
	req.Header.Set("x-api-key", s.cfg.GetTenantApiKey())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.logger.Error("request failed", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		s.logger.Error("request failed", zap.Int("status_code", resp.StatusCode))
		return "", errors.New(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("read response failed", zap.Error(err))
		return "", err
	}

	response := map[string]string{}
	if err = json.Unmarshal(data, &response); err != nil {
		s.logger.Error("unmarshal response failed",
			zap.Any("response", string(data)),
			zap.Error(err))
		return "", err
	}

	if len(response["token"]) == 0 {
		s.logger.Error("no token in response",
			zap.Any("response", string(data)))
		return "", errors.New("empty token")
	}

	return response["token"], nil
}
