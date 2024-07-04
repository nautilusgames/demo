package model

const (
	InternalAddress  = "http://demo-wallet:8080"
	CreateWalletPath = "/api/v1/wallet/create-wallet"
	TransferPath     = "/api/v1/wallet/transfer"
	GetWalletPath    = "/api/v1/wallet/get-wallet"
)

type (
	CreateWalletRequest struct {
		PlayerID int64  `json:"player_id"`
		Currency string `json:"currency"`
	}
	CreateWalletResponse struct{}
)

type (
	GetWalletRequest struct {
		PlayerID int64 `json:"player_id"`
	}
	GetWalletResponse struct {
		Data  *PlayerWallet `json:"data,omitempty"`
		Error *Error        `json:"error,omitempty"`
	}
)

type (
	BetRequest struct {
		SessionID int64 `json:"session_id,omitempty"`
		Amount    int64 `json:"amount,omitempty"`
	}
	BettResponse struct {
		Data  *Transaction `json:"data,omitempty"`
		Error *Error       `json:"error"`
	}
)

type (
	PayoutRequest struct {
		SessionID int64 `json:"session_id,omitempty"`
		Amount    int64 `json:"amount,omitempty"`
	}

	PayoutResponse struct {
		Data  *Transaction `json:"data,omitempty"`
		Error *Error       `json:"error,omitempty"`
	}
)

type (
	RefundRequest struct {
		SessionID int64 `json:"session_id,omitempty"`
		Amount    int64 `json:"amount,omitempty"`
	}
	RefundResponse struct {
		Data  *Transaction `json:"data,omitempty"`
		Error *Error       `json:"error,omitempty"`
	}
)

type (
	TransferRequest struct {
		SessionID int64  `json:"session_id"`
		GameID    string `json:"game_id"`
		PlayerID  int64  `json:"player_id"`
		Amount    int64  `json:"amount"`
	}
	TransferResponse struct {
		SessionID int64       `json:"session_id"`
		Tx        Transaction `json:"tx"`
	}
)

type Transaction struct {
	ID         int64 `json:"id,omitempty"`
	SessionID  int64 `json:"session_id,omitempty"`
	Amount     int64 `json:"amount,omitempty"`
	NewBalance int64 `json:"new_balance,omitempty"`
	CreatedAt  int64 `json:"created_at,omitempty"`
}

type PlayerWallet struct {
	Balance  int64 `json:"balance,omitempty"`
	LastTxID int64 `json:"last_tx_id,omitempty"`
}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
