package main

import (
	"fmt"
	"log"
	"net"

	"dashboard-data/grpc/server"
	"dashboard-data/pkg/config"
	"dashboard-data/pkg/logger"

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
	log.Printf("gRPC Server started on %s", addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	logger.InitLogger()

	// 启动gRPC服务器
	log.Printf("Starting Dashboard Data Service gRPC server on port %d", cfg.GRPCPort)
	startGRPCServer(cfg)
}
