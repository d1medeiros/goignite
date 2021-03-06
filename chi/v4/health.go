package gichi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/b2wdigital/goignite/rest/response"
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(ctx context.Context) http.HandlerFunc {
	resp, httpCode := response.NewHealth(ctx)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		json.NewEncoder(w).Encode(resp)
	}
}
