package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"micro-go/internal/api"
	"micro-go/internal/mock"
	"micro-go/pkg/config"
	"micro-go/pkg/logger"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// 初始化日志
	log := logger.New()
	log.Info("Starting server", map[string]interface{}{"host": cfg.Server.Host, "port": cfg.Server.Port})

	// 初始化 Mock 数据
	userMock := mock.NewUserMock()

	// 初始化 API 层和路由
	router := api.SetupRouter(userMock, log)

	// 启动服务器
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Info("Server started", map[string]interface{}{"address": serverAddr})

	// 优雅关闭
	go func() {
		if err := router.Run(serverAddr); err != nil {
			log.Fatal("Failed to start server", map[string]interface{}{"error": err.Error()})
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server", nil)
}
