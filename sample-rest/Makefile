# 构建和运行脚本

# 项目名称
APP_NAME := micro-go

# 构建目录
BUILD_DIR := ./build

# 源码目录
SRC_DIR := ./cmd/server

# 依赖管理
deps:
	@echo "Installing dependencies..."
	go mod tidy

# 构建项目
build:
	@echo "Building project..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)/main.go

# 运行项目
run:
	@echo "Running project..."
	go run $(SRC_DIR)/main.go

# 清理构建文件
clean:
	@echo "Cleaning build files..."
	rm -rf $(BUILD_DIR)

# 帮助信息
help:
	@echo "Available commands:"
	@echo "  make deps     - Install dependencies"
	@echo "  make build    - Build project"
	@echo "  make run      - Run project"
	@echo "  make clean    - Clean build files"
	@echo "  make help     - Show this help"

.PHONY: deps build run clean help
