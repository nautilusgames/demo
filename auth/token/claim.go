package token

import "github.com/golang-jwt/jwt/v5"

type MyClaim struct {
	Object   string `json:"object"`
	PlayerID int64  `json:"player_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
