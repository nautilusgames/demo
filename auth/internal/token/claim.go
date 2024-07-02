package token

import "github.com/golang-jwt/jwt/v5"

type MyClaim struct {
	Username string `json:"username"`
	PlayerId int64  `json:"player_id"`
	jwt.RegisteredClaims
}
