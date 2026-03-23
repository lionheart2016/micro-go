package api

import (
	"dashboard-rest/grpc/client"
	"dashboard-rest/model"
	"dashboard-rest/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

// 客户分析相关接口
func (h *DashboardHandler) GetCustomerDistribution(c *gin.Context) {
	h.logger.Info("Getting customer distribution")

	dimension := c.Query("dimension")
	if dimension == "" {
		dimension = "age"
	}

	response, err := h.client.GetCustomerDistribution(c.Request.Context(), dimension)
	if err != nil {
		h.logger.Error("Failed to get customer distribution: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customer distribution"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetCustomerBehavior(c *gin.Context) {
	h.logger.Info("Getting customer behavior")

	behaviorType := c.Query("type")
	if behaviorType == "" {
		behaviorType = "retention"
	}

	response, err := h.client.GetCustomerBehavior(c.Request.Context(), behaviorType)
	if err != nil {
		h.logger.Error("Failed to get customer behavior: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customer behavior"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// 交易分析相关接口
func (h *DashboardHandler) GetStockTrade(c *gin.Context) {
	h.logger.Info("Getting stock trade")

	tradeType := c.Query("type")
	if tradeType == "" {
		tradeType = "execution"
	}

	response, err := h.client.GetStockTrade(c.Request.Context(), tradeType)
	if err != nil {
		h.logger.Error("Failed to get stock trade: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stock trade"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetFundTrade(c *gin.Context) {
	h.logger.Info("Getting fund trade")

	tradeType := c.Query("type")
	if tradeType == "" {
		tradeType = "销售情况"
	}

	response, err := h.client.GetFundTrade(c.Request.Context(), tradeType)
	if err != nil {
		h.logger.Error("Failed to get fund trade: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get fund trade"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PI用户分析相关接口
func (h *DashboardHandler) GetPIBasic(c *gin.Context) {
	h.logger.Info("Getting PI basic information")

	response, err := h.client.GetPIBasic(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get PI basic information: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get PI basic information"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetPIBehavior(c *gin.Context) {
	h.logger.Info("Getting PI behavior")

	behaviorType := c.Query("type")
	if behaviorType == "" {
		behaviorType = "stock"
	}

	response, err := h.client.GetPIBehavior(c.Request.Context(), behaviorType)
	if err != nil {
		h.logger.Error("Failed to get PI behavior: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get PI behavior"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// 开户主题相关接口
func (h *DashboardHandler) GetAccountProcess(c *gin.Context) {
	h.logger.Info("Getting account process")

	response, err := h.client.GetAccountProcess(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get account process: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get account process"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetAccountException(c *gin.Context) {
	h.logger.Info("Getting account exception")

	response, err := h.client.GetAccountException(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get account exception: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get account exception"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// IPO主题相关接口
func (h *DashboardHandler) GetIPOSubscription(c *gin.Context) {
	h.logger.Info("Getting IPO subscription")

	response, err := h.client.GetIPOSubscription(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get IPO subscription: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get IPO subscription"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetIPOAnaalysis(c *gin.Context) {
	h.logger.Info("Getting IPO analysis")

	response, err := h.client.GetIPOAnaalysis(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get IPO analysis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get IPO analysis"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// 融资主题相关接口
func (h *DashboardHandler) GetFinanceCustomer(c *gin.Context) {
	h.logger.Info("Getting finance customer")

	response, err := h.client.GetFinanceCustomer(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get finance customer: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get finance customer"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetFinanceStock(c *gin.Context) {
	h.logger.Info("Getting finance stock")

	response, err := h.client.GetFinanceStock(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get finance stock: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get finance stock"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// 数据钻取相关接口
func (h *DashboardHandler) GetDrilldownDetail(c *gin.Context) {
	h.logger.Info("Getting drilldown detail")

	dataType := c.Query("type")
	if dataType == "" {
		dataType = "customer"
	}

	id := c.Query("id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	response, err := h.client.GetDrilldownDetail(c.Request.Context(), dataType, id, page, pageSize)
	if err != nil {
		h.logger.Error("Failed to get drilldown detail: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get drilldown detail"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetDrilldownTrend(c *gin.Context) {
	h.logger.Info("Getting drilldown trend")

	dataType := c.Query("type")
	if dataType == "" {
		dataType = "indicator"
	}

	id := c.Query("id")
	period := c.DefaultQuery("period", "1m")

	response, err := h.client.GetDrilldownTrend(c.Request.Context(), dataType, id, period)
	if err != nil {
		h.logger.Error("Failed to get drilldown trend: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get drilldown trend"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// 认证相关接口
func (h *DashboardHandler) Login(c *gin.Context) {
	h.logger.Info("User login")

	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid login request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login request"})
		return
	}

	response, err := h.client.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		h.logger.Error("Login failed: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetAuthInfo(c *gin.Context) {
	h.logger.Info("Getting auth info")

	response, err := h.client.GetAuthInfo(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get auth info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get auth info"})
		return
	}

	c.JSON(http.StatusOK, response)
}
