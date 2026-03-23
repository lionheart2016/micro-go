package api

import (
	"dashboard-rest/grpc/client"
	"dashboard-rest/pkg/logger"
	"github.com/gin-gonic/gin"
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

func (h *DashboardHandler) GetCoreIndicators(c *gin.Context) {
	h.logger.Info("Getting core indicators")

	period := c.Query("period")
	if period == "" {
		period = "7d"
	}

	response, err := h.client.GetCoreIndicators(c.Request.Context(), period)
	if err != nil {
		h.logger.Error("Failed to get core indicators: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get core indicators"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetPerformance(c *gin.Context) {
	h.logger.Info("Getting performance")

	period := c.Query("period")
	if period == "" {
		period = "7d"
	}

	response, err := h.client.GetPerformance(c.Request.Context(), period)
	if err != nil {
		h.logger.Error("Failed to get performance: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get performance"})
		return
	}

	c.JSON(http.StatusOK, response)
}
