package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dashboard-data/internal/repository"
	"dashboard-data/pkg/config"
)

// HTTP服务器
func startHTTPServer(cfg *config.Config) {
	// 创建路由
	http.HandleFunc("/api/dashboard/core-indicators", getCoreIndicators)
	http.HandleFunc("/api/dashboard/performance", getPerformance)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.HTTPPort)
	log.Printf("HTTP Server started on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

// 响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// getCoreIndicators 获取核心指标
func getCoreIndicators(w http.ResponseWriter, r *http.Request) {
	dashboardRepo := repository.NewDashboardRepository()
	indicators, err := dashboardRepo.GetCoreIndicators()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    indicators,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getPerformance 获取业绩表现
func getPerformance(w http.ResponseWriter, r *http.Request) {
	dashboardRepo := repository.NewDashboardRepository()
	performance, err := dashboardRepo.GetPerformance()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    performance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
