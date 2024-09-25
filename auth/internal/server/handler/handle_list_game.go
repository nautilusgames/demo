package handler

import (
	"net/http"

	"github.com/carlmjohnson/requests"
	"github.com/nautilusgames/demo/config/pb"
	sgbuilder "github.com/nautilusgames/sdk-go/builder"
	"go.uber.org/zap"
)

func (h *Handler) HandleListGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := h.authorize(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var tenantInfo *pb.TenantInfo
		for _, v := range h.cfg.GetTenants() {
			if v.GetTenantId() == payload.Object {
				tenantInfo = v
				break
			}
		}
		if tenantInfo == nil {
			http.Error(w, "tenant not found", http.StatusInternalServerError)
			return
		}

		var resp interface{}
		err = requests.
			URL(tenantInfo.ListGameUrl).
			Header("x-tenant-id", tenantInfo.TenantId).
			Header("x-api-key", tenantInfo.TenantApiKey).
			ToJSON(&resp).
			Fetch(r.Context())
		if err != nil {
			h.logger.Error("failed to fetch list game", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sgbuilder.SendReply(w, resp)
	}
}
