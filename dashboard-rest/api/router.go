package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *DashboardHandler) *gin.Engine {
	router := gin.Default()

	// 核心指标相关路由
	router.GET("/api/core-indicators", handler.GetCoreIndicators)

	// 业绩表现相关路由
	router.GET("/api/performance", handler.GetPerformance)

	return router
}
