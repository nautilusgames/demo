package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

const (
	InternalAddress  = "http://demo-wallet:8080"
	CreateWalletPath = "/wallet/create-wallet"
)

type CreateWalletRequest struct {
	PlayerID int64  `json:"player_id"`
	Currency string `json:"currency"`
}

type CreateWalletResponse struct{}

func CreateWallet(logger *zap.Logger, request *CreateWalletRequest) (*CreateWalletResponse, error) {
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		logger.Error("failed to encode body", zap.Error(err))
		return nil, err
	}

	url := fmt.Sprintf("%s%s", InternalAddress, CreateWalletPath)
	resp, err := http.Post(url, "application/json", &body)
	if err != nil {
		logger.Error("failed to post http request", zap.Error(err))
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logger.Error("failed to create wallet", zap.Any("status", resp.StatusCode))
		return nil, errors.New("failed to create wallet")
	}

	return &CreateWalletResponse{}, nil
}
