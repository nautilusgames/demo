package model

type SignUpRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
	Currency    string `json:"currency"`
}
type SignUpResponse struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Token       string `json:"token"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type SignInResponse struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Token       string `json:"token"`
	Currency    string `json:"currency"`
}

type VerifyPlayerRequest struct {
	Token string `json:"token"`
}
type VerifyPlayerResponse struct {
	Data *PlayerInfo `json:"data"`
}

type PlayerInfo struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type CreateSessionRequest struct {
	GameId string `json:"game_id"`
}
type CreateSessionResponse struct {
	TenantId int64  `json:"tenant_id"`
	Token    string `json:"token"`
}
