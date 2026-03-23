package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *DashboardHandler) *gin.Engine {
	router := gin.Default()

	// 核心指标相关路由
	router.GET("/api/dashboard/core-indicators", handler.GetCoreIndicators)
	router.GET("/api/dashboard/performance", handler.GetPerformance)

	// 客户分析相关路由
	router.GET("/api/customer/distribution", handler.GetCustomerDistribution)
	router.GET("/api/customer/behavior", handler.GetCustomerBehavior)

	// 交易分析相关路由
	router.GET("/api/trade/stock", handler.GetStockTrade)
	router.GET("/api/trade/fund", handler.GetFundTrade)

	// PI用户分析相关路由
	router.GET("/api/pi/basic", handler.GetPIBasic)
	router.GET("/api/pi/behavior", handler.GetPIBehavior)

	// 开户主题相关路由
	router.GET("/api/account/process", handler.GetAccountProcess)
	router.GET("/api/account/exception", handler.GetAccountException)

	// IPO主题相关路由
	router.GET("/api/ipo/subscription", handler.GetIPOSubscription)
	router.GET("/api/ipo/analysis", handler.GetIPOAnaalysis)

	// 融资主题相关路由
	router.GET("/api/finance/customer", handler.GetFinanceCustomer)
	router.GET("/api/finance/stock", handler.GetFinanceStock)

	// 数据钻取相关路由
	router.GET("/api/drilldown/detail", handler.GetDrilldownDetail)
	router.GET("/api/drilldown/trend", handler.GetDrilldownTrend)

	// 认证相关路由
	router.POST("/api/auth/login", handler.Login)
	router.GET("/api/auth/info", handler.GetAuthInfo)

	return router
}
