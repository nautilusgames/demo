package model

const (
	InternalAddress = "http://demo-wallet:8080"
	CreatePath      = "/wallet/create"
	BetPath         = "/wallet/bet"
	PayoutPath      = "/wallet/payout"
	RefundPath      = "/wallet/refund"
	GetPath         = "/wallet/get"
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
		GameID    string `json:"game_id"`
		PlayerID  int64  `json:"player_id"`
		SessionID int64  `json:"session_id,omitempty"`
		Amount    int64  `json:"amount,omitempty"`
	}
	PayoutRequest struct {
		GameID    string `json:"game_id"`
		PlayerID  int64  `json:"player_id"`
		SessionID int64  `json:"session_id,omitempty"`
		Amount    int64  `json:"amount,omitempty"`
	}
	RefundRequest struct {
		GameID    string `json:"game_id"`
		PlayerID  int64  `json:"player_id"`
		SessionID int64  `json:"session_id,omitempty"`
		Amount    int64  `json:"amount,omitempty"`
	}
	RollbackRequest struct {
		GameID    string `json:"game_id"`
		PlayerID  int64  `json:"player_id"`
		SessionID int64  `json:"session_id,omitempty"`
		Amount    int64  `json:"amount,omitempty"`
	}
	Response struct {
		Data  *Transaction `json:"data,omitempty"`
		Error *Error       `json:"error"`
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
