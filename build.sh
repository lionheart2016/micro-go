#!/bin/bash

# 编译脚本，为所有 Golang 服务生成 amd64 架构的可执行文件

set -e

# 定义颜色
GREEN="\033[0;32m"
YELLOW="\033[1;33m"
RED="\033[0;31m"
NC="\033[0m" # No Color

# 打印带颜色的消息
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 编译单个服务
build_service() {
    local service_name=$1
    local service_dir=$2
    
    log_info "开始编译 $service_name 服务..."
    
    cd "$service_dir"
    
    # 创建输出目录
    mkdir -p bin
    
    # 清理旧文件
    rm -f bin/server
    
    # 编译
    go build -o bin/server main.go
    
    if [ $? -eq 0 ]; then
        log_info "$service_name 服务编译成功"
    else
        log_error "$service_name 服务编译失败"
        return 1
    fi
    
    cd ..
    return 0
}

# 主函数
main() {
    log_info "开始编译所有 Golang 服务..."
    
    # 设置编译环境变量，生成 amd64 架构的可执行文件
    export GOOS=linux
    export GOARCH=amd64
    export CGO_ENABLED=0
    
    # 编译服务列表
    services=("dashboard-data" "dashboard-rest" "sample-rest")
    
    # 并行编译
    for service in "${services[@]}"; do
        build_service "$service" "$service"
    done
    
    log_info "所有 Golang 服务编译完成！"
}

# 执行主函数
main
