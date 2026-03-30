package api

import (
	"dashboard-rest/grpc/client"
	"dashboard-rest/model"
	"dashboard-rest/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type DashboardHandler struct {
	client *client.DashboardClient
}

func NewDashboardHandler(client *client.DashboardClient) *DashboardHandler {
	return &DashboardHandler{
		client: client,
	}
}

func (h *DashboardHandler) GetCoreIndicators(c *gin.Context) {
	logger.Info("Getting core indicators")

	period := c.Query("period")
	if period == "" {
		period = "7d"
	}

	response, err := h.client.GetCoreIndicators(c.Request.Context(), period)
	if err != nil {
		logger.Error("Failed to get core indicators",
			zap.Error(err),
			zap.String("period", period),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get core indicators"})
		return
	}
	
	logger.Info("Core indicators retrieved successfully",
		zap.String("period", period),
	)

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetPerformance(c *gin.Context) {
	logger.Info("Getting performance")

	period := c.Query("period")
	if period == "" {
		period = "7d"
	}

	response, err := h.client.GetPerformance(c.Request.Context(), period)
	if err != nil {
		logger.Error("Failed to get performance",
			zap.Error(err),
			zap.String("period", period),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get performance"})
		return
	}

	logger.Info("Performance retrieved successfully",
		zap.String("period", period),
	)

	c.JSON(http.StatusOK, response)
}

// 客户分析相关接口
func (h *DashboardHandler) GetCustomerDistribution(c *gin.Context) {
	logger.Info("Getting customer distribution")

	dimension := c.Query("dimension")
	if dimension == "" {
		dimension = "age"
	}

	response, err := h.client.GetCustomerDistribution(c.Request.Context(), dimension)
	if err != nil {
		logger.Error("Failed to get customer distribution",
			zap.Error(err),
			zap.String("dimension", dimension),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customer distribution"})
		return
	}

	logger.Info("Customer distribution retrieved successfully",
		zap.String("dimension", dimension),
	)

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetCustomerBehavior(c *gin.Context) {
	logger.Info("Getting customer behavior")

	behaviorType := c.Query("type")
	if behaviorType == "" {
		behaviorType = "retention"
	}

	response, err := h.client.GetCustomerBehavior(c.Request.Context(), behaviorType)
	if err != nil {
		logger.Error("Failed to get customer behavior",
			zap.Error(err),
			zap.String("type", behaviorType),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customer behavior"})
		return
	}

	logger.Info("Customer behavior retrieved successfully",
		zap.String("type", behaviorType),
	)

	c.JSON(http.StatusOK, response)
}

// 交易分析相关接口
func (h *DashboardHandler) GetStockTrade(c *gin.Context) {
	logger.Info("Getting stock trade")

	tradeType := c.Query("type")
	if tradeType == "" {
		tradeType = "execution"
	}

	response, err := h.client.GetStockTrade(c.Request.Context(), tradeType)
	if err != nil {
		logger.Error("Failed to get stock trade",
			zap.Error(err),
			zap.String("type", tradeType),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stock trade"})
		return
	}

	logger.Info("Stock trade retrieved successfully",
		zap.String("type", tradeType),
	)

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetFundTrade(c *gin.Context) {
	logger.Info("Getting fund trade")

	tradeType := c.Query("type")
	if tradeType == "" {
		tradeType = "销售情况"
	}

	response, err := h.client.GetFundTrade(c.Request.Context(), tradeType)
	if err != nil {
		logger.Error("Failed to get fund trade",
			zap.Error(err),
			zap.String("type", tradeType),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get fund trade"})
		return
	}

	logger.Info("Fund trade retrieved successfully",
		zap.String("type", tradeType),
	)

	c.JSON(http.StatusOK, response)
}

// PI用户分析相关接口
func (h *DashboardHandler) GetPIBasic(c *gin.Context) {
	logger.Info("Getting PI basic information")

	response, err := h.client.GetPIBasic(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get PI basic information",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get PI basic information"})
		return
	}

	logger.Info("PI basic information retrieved successfully")

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetPIBehavior(c *gin.Context) {
	logger.Info("Getting PI behavior")

	behaviorType := c.Query("type")
	if behaviorType == "" {
		behaviorType = "stock"
	}

	response, err := h.client.GetPIBehavior(c.Request.Context(), behaviorType)
	if err != nil {
		logger.Error("Failed to get PI behavior",
			zap.Error(err),
			zap.String("type", behaviorType),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get PI behavior"})
		return
	}

	logger.Info("PI behavior retrieved successfully",
		zap.String("type", behaviorType),
	)

	c.JSON(http.StatusOK, response)
}

// 开户主题相关接口
func (h *DashboardHandler) GetAccountProcess(c *gin.Context) {
	logger.Info("Getting account process")

	response, err := h.client.GetAccountProcess(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get account process",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get account process"})
		return
	}

	logger.Info("Account process retrieved successfully")

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetAccountException(c *gin.Context) {
	logger.Info("Getting account exception")

	response, err := h.client.GetAccountException(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get account exception",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get account exception"})
		return
	}

	logger.Info("Account exception retrieved successfully")

	c.JSON(http.StatusOK, response)
}

// IPO主题相关接口
func (h *DashboardHandler) GetIPOSubscription(c *gin.Context) {
	logger.Info("Getting IPO subscription")

	response, err := h.client.GetIPOSubscription(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get IPO subscription",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get IPO subscription"})
		return
	}

	logger.Info("IPO subscription retrieved successfully")

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetIPOAnaalysis(c *gin.Context) {
	logger.Info("Getting IPO analysis")

	response, err := h.client.GetIPOAnaalysis(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get IPO analysis",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get IPO analysis"})
		return
	}

	logger.Info("IPO analysis retrieved successfully")

	c.JSON(http.StatusOK, response)
}

// 融资主题相关接口
func (h *DashboardHandler) GetFinanceCustomer(c *gin.Context) {
	logger.Info("Getting finance customer")

	response, err := h.client.GetFinanceCustomer(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get finance customer",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get finance customer"})
		return
	}

	logger.Info("Finance customer retrieved successfully")

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetFinanceStock(c *gin.Context) {
	logger.Info("Getting finance stock")

	response, err := h.client.GetFinanceStock(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get finance stock",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get finance stock"})
		return
	}

	logger.Info("Finance stock retrieved successfully")

	c.JSON(http.StatusOK, response)
}

// 数据钻取相关接口
func (h *DashboardHandler) GetDrilldownDetail(c *gin.Context) {
	logger.Info("Getting drilldown detail")

	dataType := c.Query("type")
	if dataType == "" {
		dataType = "customer"
	}

	id := c.Query("id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	response, err := h.client.GetDrilldownDetail(c.Request.Context(), dataType, id, page, pageSize)
	if err != nil {
		logger.Error("Failed to get drilldown detail",
			zap.Error(err),
			zap.String("type", dataType),
			zap.String("id", id),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get drilldown detail"})
		return
	}

	logger.Info("Drilldown detail retrieved successfully",
		zap.String("type", dataType),
		zap.String("id", id),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
	)

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetDrilldownTrend(c *gin.Context) {
	logger.Info("Getting drilldown trend")

	dataType := c.Query("type")
	if dataType == "" {
		dataType = "indicator"
	}

	id := c.Query("id")
	period := c.DefaultQuery("period", "1m")

	response, err := h.client.GetDrilldownTrend(c.Request.Context(), dataType, id, period)
	if err != nil {
		logger.Error("Failed to get drilldown trend",
			zap.Error(err),
			zap.String("type", dataType),
			zap.String("id", id),
			zap.String("period", period),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get drilldown trend"})
		return
	}

	logger.Info("Drilldown trend retrieved successfully",
		zap.String("type", dataType),
		zap.String("id", id),
		zap.String("period", period),
	)

	c.JSON(http.StatusOK, response)
}

// 认证相关接口
func (h *DashboardHandler) Login(c *gin.Context) {
	logger.Info("User login")

	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid login request",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login request"})
		return
	}

	response, err := h.client.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		logger.Error("Login failed",
			zap.Error(err),
			zap.String("username", req.Username),
		)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	logger.Info("Login successful",
		zap.String("username", req.Username),
	)

	c.JSON(http.StatusOK, response)
}

func (h *DashboardHandler) GetAuthInfo(c *gin.Context) {
	logger.Info("Getting auth info")

	response, err := h.client.GetAuthInfo(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get auth info",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get auth info"})
		return
	}

	logger.Info("Auth info retrieved successfully")

	c.JSON(http.StatusOK, response)
}
