package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"micro-go/api"
	"micro-go/pkg/config"
	"micro-go/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		// 这里使用标准库 fmt 是因为日志组件还未初始化
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// 初始化日志
	logger.InitLogger("sample-rest")
	defer logger.Sync()
	logger.Info("Starting server",
		zap.String("host", cfg.Server.Host),
		zap.Int("port", cfg.Server.Port),
	)

	// 初始化 API 层和路由
	router := api.SetupRouter()

	// 启动服务器
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	logger.Info("Server started",
		zap.String("address", serverAddr),
	)

	// 优雅关闭
	go func() {
		if err := router.Run(serverAddr); err != nil {
			logger.Fatal("Failed to start server",
				zap.Error(err),
				zap.String("address", serverAddr),
			)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server")
}
