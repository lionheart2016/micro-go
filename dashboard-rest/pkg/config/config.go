package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort int
	GRPCHost   string
	GRPCPort   int
}

func Load() *Config {
	port, _ := strconv.Atoi(getEnv("SERVER_PORT", "8082"))
	grpcPort, _ := strconv.Atoi(getEnv("GRPC_PORT", "50051"))

	return &Config{
		ServerPort: port,
		GRPCHost:   getEnv("GRPC_HOST", "localhost"),
		GRPCPort:   grpcPort,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
