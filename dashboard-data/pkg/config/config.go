package config

import (
	"os"
	"strconv"
)

// Config 配置结构体
type Config struct {
	GRPCPort int
}

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	grpcPort, err := strconv.Atoi(getEnv("GRPC_PORT", "50051"))
	if err != nil {
		grpcPort = 50051
	}

	return &Config{
		GRPCPort: grpcPort,
	}, nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
