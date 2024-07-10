package handler

import "net/http"

func (h *Handler) HandleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
