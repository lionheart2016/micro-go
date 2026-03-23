package api

import (
	"net/http"
)

func SetupRouter(handler *DashboardHandler) http.Handler {
	mux := http.NewServeMux()

	// 核心指标相关路由
	mux.HandleFunc("/api/core-indicators", handler.GetCoreIndicators)

	// 业绩表现相关路由
	mux.HandleFunc("/api/performance", handler.GetPerformance)

	return mux
}
