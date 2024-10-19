package controller

import (
	"encoding/json"
	"golang.org/x/exp/slog"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

type IHealthController interface {
	health()
}

type HealthController struct {
	Logger *slog.Logger
}

func (h *HealthController) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Health{Status: "OK"}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.Logger.Error("Failed on marshall health response", "error", err)
		http.Error(w, "Failed on marshall health response", http.StatusInternalServerError)
		return
	}
}
