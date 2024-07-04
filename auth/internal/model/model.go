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
}

type VerifyPlayerRequest struct {
	Token string `json:"token"`
}
type VerifyPlayerResponse struct {
	PlayerId    int64  `json:"player_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

type (
	CreateTenantTokenRequest  struct{}
	CreateTenantTokenResponse struct {
		TenantId int64  `json:"tenant_id"`
		Token    string `json:"token"`
	}
)
