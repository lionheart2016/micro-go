package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config 应用配置结构
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int
	Host string
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	DSN string
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_PORT: %w", err)
	}

	redisDB, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_DB: %w", err)
	}

	return &Config{
		Server: ServerConfig{
			Port: port,
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
		},
		Database: DatabaseConfig{
			DSN: getEnv("DATABASE_DSN", ""),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       redisDB,
		},
	}, nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
