package main

import (
	"dashboard-rest/api"
	"dashboard-rest/grpc/client"
	"dashboard-rest/pkg/config"
	"dashboard-rest/pkg/logger"
	"fmt"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化日志
	log := logger.New()
	log.Info("Starting dashboard-rest service")

	// 初始化gRPC客户端
	dashboardClient, err := client.New(cfg.GRPCHost, cfg.GRPCPort, log)
	if err != nil {
		log.Error("Failed to initialize gRPC client: %v", err)
		return
	}

	// 初始化API处理器
	handler := api.NewDashboardHandler(dashboardClient, log)

	// 设置路由
	router := api.SetupRouter(handler)

	// 启动服务器
	serverAddr := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Info("Server starting on %s", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		log.Error("Failed to start server on %s: %v", serverAddr, err)
		return
	}
}
