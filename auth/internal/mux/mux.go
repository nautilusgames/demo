package mux

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func New(logger *zap.Logger) *http.ServeMux {
	// Flag gets printed as a page
	mux := http.NewServeMux()
	// Health endpoint
	mux.HandleFunc("/status", httpHealth())
	mux.HandleFunc("/api/v1/player/verify", httpAuth(logger))

	return mux
}

func httpAuth(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"status\": \"ok\"}")
	}
}

func httpHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
