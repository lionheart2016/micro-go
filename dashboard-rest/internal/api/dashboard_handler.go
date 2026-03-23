package api

import (
	"dashboard-rest/internal/grpc/client"
	"dashboard-rest/pkg/logger"
	"encoding/json"
	"net/http"
)

type DashboardHandler struct {
	client *client.DashboardClient
	logger *logger.Logger
}

func NewDashboardHandler(client *client.DashboardClient, logger *logger.Logger) *DashboardHandler {
	return &DashboardHandler{
		client: client,
		logger: logger,
	}
}

func (h *DashboardHandler) GetCoreIndicators(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Getting core indicators")

	period := r.URL.Query().Get("period")
	if period == "" {
		period = "7d"
	}

	response, err := h.client.GetCoreIndicators(r.Context(), period)
	if err != nil {
		h.logger.Error("Failed to get core indicators: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *DashboardHandler) GetPerformance(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Getting performance")

	period := r.URL.Query().Get("period")
	if period == "" {
		period = "7d"
	}

	response, err := h.client.GetPerformance(r.Context(), period)
	if err != nil {
		h.logger.Error("Failed to get performance: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
