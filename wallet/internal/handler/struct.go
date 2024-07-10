package handler

type CreateWalletRequest struct {
	PlayerID int64  `json:"player_id"`
	Currency string `json:"currency"`
}

type CreateWalletResponse struct{}
