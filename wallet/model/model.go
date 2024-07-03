package model

// Get wallet
type GetWalletRequest struct {
	PlayerID int64 `json:"player_id"`
}
type GetWalletResponse struct {
	Balance int64 `json:"balance"`
}

// Create player wallet
type CreateWalletRequest struct {
	PlayerID int64  `json:"player_id"`
	Currency string `json:"currency"`
}
type CreateWalletResponse struct{}

// Transfer
type TransferRequest struct {
	SessionID int64  `json:"session_id"`
	GameID    string `json:"game_id"`
	PlayerID  int64  `json:"player_id"`
	Amount    int64  `json:"amount"`
}
type TransferResponse struct {
	SessionID int64       `json:"session_id"`
	Tx        Transaction `json:"tx"`
}
type Transaction struct {
	TxID      int64  `json:"tx_id"`
	PlayerID  int64  `json:"player_id"`
	GameID    string `json:"game_id"`
	Amount    int64  `json:"amount"`
	CreatedAt int64  `json:"created_at"`
}
