package main

import (
	"fmt"
	"net"

	"dashboard-data/grpc/server"
	"dashboard-data/pkg/config"
	"dashboard-data/pkg/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// startGRPCServer 启动gRPC服务器
func startGRPCServer(cfg *config.Config) {
	// 创建gRPC服务器
	s := grpc.NewServer()

	// 注册服务
	server.RegisterDataService(s)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.GRPCPort)
	logger.Info("gRPC Server started",
		zap.String("address", addr),
	)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal("Failed to listen",
			zap.Error(err),
		)
	}

	if err := s.Serve(lis); err != nil {
		logger.Fatal("Failed to serve",
			zap.Error(err),
		)
	}
}

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		// 这里使用标准库 log 是因为日志组件还未初始化
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// 初始化日志
	logger.InitLogger("dashboard-data")
	defer logger.Sync()

	// 启动gRPC服务器
	logger.Info("Starting Dashboard Data Service gRPC server",
		zap.Int("port", cfg.GRPCPort),
	)
	startGRPCServer(cfg)
}
