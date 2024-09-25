package verifier

import (
	"github.com/nautilusgames/sdk-go/webhook"

	"github.com/nautilusgames/demo/auth/token"
	"github.com/nautilusgames/demo/config/pb"
)

func Verify(cfg *pb.Config, token token.Maker, header *webhook.HookRequestHeader) (*token.Payload, *webhook.Error) {
	// validate headers
	if header.XTenantId == "" {
		return nil, &webhook.Error{Code: webhook.ErrInvalidTenantID, Message: "empty tenant_id"}
	}
	if header.XApiKey == "" {
		return nil, &webhook.Error{Code: webhook.ErrInvalidTenantApiKey, Message: "empty api_key"}
	}
	if header.XGameId == "" {
		return nil, &webhook.Error{Code: webhook.ErrInvalidGameID, Message: "empty game_id"}
	}
	if header.XTenantPlayerToken == "" {
		return nil, &webhook.Error{Code: webhook.ErrEmptyPlayerSession, Message: "empty tenant_player_token"}
	}

	var tenantInfo *pb.TenantInfo
	for _, v := range cfg.GetTenants() {
		if v.GetTenantId() == header.XTenantId {
			tenantInfo = v
			break
		}
	}
	if tenantInfo == nil {
		return nil, &webhook.Error{Code: webhook.ErrInvalidTenantID, Message: "invalid tenant_id"}
	}
	if header.XApiKey != tenantInfo.GetTenantApiKey() {
		return nil, &webhook.Error{Code: webhook.ErrInvalidTenantApiKey, Message: "invalid api_key"}
	}

	// validate tenant token
	payload, err := token.VerifyToken(header.XTenantPlayerToken)
	if err != nil {
		return nil, &webhook.Error{Code: webhook.ErrInvalidPlayerSession, Message: "invalid tenant_player_token"}
	}

	if payload.Object != header.XGameId {
		return nil, &webhook.Error{Code: webhook.ErrInvalidGameID, Message: "invalid game_id"}
	}

	return payload, nil
}
