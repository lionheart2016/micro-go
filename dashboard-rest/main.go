package main

import (
	"dashboard-rest/api"
	"dashboard-rest/grpc/client"
	"dashboard-rest/pkg/config"
	"dashboard-rest/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化日志
	logger.InitLogger("dashboard-rest")
	defer logger.Sync()
	logger.Info("Starting dashboard-rest service")

	// 初始化gRPC客户端
	dashboardClient, err := client.New(cfg.GRPCHost, cfg.GRPCPort)
	if err != nil {
		logger.Error("Failed to initialize gRPC client",
			zap.Error(err),
			zap.String("host", cfg.GRPCHost),
			zap.Int("port", cfg.GRPCPort),
		)
		return
	}

	// 初始化API处理器
	handler := api.NewDashboardHandler(dashboardClient)

	// 设置路由
	router := api.SetupRouter(handler)

	// 启动服务器
	serverAddr := fmt.Sprintf(":%d", cfg.ServerPort)
	logger.Info("Server starting",
		zap.String("address", serverAddr),
	)

	if err := router.Run(serverAddr); err != nil {
		logger.Error("Failed to start server",
			zap.Error(err),
			zap.String("address", serverAddr),
		)
		return
	}
}
