#!/bin/bash

# Docker 构建脚本
# 用于批量构建所有服务的 Docker 镜像

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 项目根目录
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# 默认版本
DEFAULT_VERSION="latest"

# 服务列表
SERVICES="dashboard-rest dashboard-data sample-rest"

# 打印信息
print_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 显示帮助信息
show_help() {
    cat << EOF
Docker 镜像构建脚本

用法：$0 [选项]

选项:
    -s, --service NAME    构建指定服务 (默认：所有服务)
    -v, --version VER     设置镜像版本 (默认：latest)
    -p, --push            构建后推送到镜像仓库
    -n, --no-cache        不使用缓存
    -h, --help            显示此帮助信息

可用服务:
    $(for svc in ${SERVICES}; do echo "    - ${svc}"; done)

示例:
    $0                          # 构建所有服务
    $0 -s dashboard-rest        # 构建 dashboard-rest 服务
    $0 -v 1.0.0                 # 使用版本 1.0.0 构建所有服务
    $0 -s dashboard-rest -v 1.0.0  # 使用版本 1.0.0 构建 dashboard-rest 服务

EOF
}

# 本地编译服务
build_local() {
    local service=$1
    local service_dir="${PROJECT_ROOT}/${service}"
    
    print_info "本地编译 ${service} 服务..."
    
    # 检查服务目录是否存在
    if [ ! -d "${service_dir}" ]; then
        print_error "服务目录不存在：${service_dir}"
        return 1
    fi
    
    # 检查 Go 是否安装
    if ! command -v go &> /dev/null; then
        print_error "Go 未安装或未在 PATH 中"
        return 1
    fi
    
    # 创建 bin 目录
    mkdir -p "${service_dir}/bin"
    
    # 进入服务目录
    cd "${service_dir}"
    
    # 下载依赖
    print_info "下载依赖..."
    go mod tidy
    
    # 根据服务选择正确的入口点
    local main_path="."
    if [ "${service}" = "sample-rest" ]; then
        main_path="./cmd/server"
    fi
    
    # 编译二进制文件
    print_info "编译二进制文件..."
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server "${main_path}"
    
    # 检查编译是否成功
    if [ ! -f "bin/server" ]; then
        print_error "编译失败！"
        cd "${PROJECT_ROOT}"
        return 1
    fi
    
    print_info "${service} 本地编译成功！"
    
    # 返回项目根目录
    cd "${PROJECT_ROOT}"
    
    return 0
}

# 构建单个服务
build_service() {
    local service=$1
    local version=$2
    local no_cache=$3
    
    local service_dir="${PROJECT_ROOT}/${service}"
    local dockerfile="${service_dir}/Dockerfile"
    local binary_path="${service_dir}/bin/server"
    
    # 检查服务目录是否存在
    if [ ! -d "${service_dir}" ]; then
        print_error "服务目录不存在：${service_dir}"
        return 1
    fi
    
    # 检查 Dockerfile 是否存在
    if [ ! -f "${dockerfile}" ]; then
        print_error "Dockerfile 不存在：${dockerfile}"
        return 1
    fi
    
    # 检查二进制文件是否存在
    if [ ! -f "${binary_path}" ]; then
        print_error "二进制文件不存在：${binary_path}"
        print_error "请先进行本地编译"
        return 1
    fi
    
    # 构建镜像标签
    local image_name="micro-go/${service}"
    local image_tags=()
    
    # 添加版本标签
    image_tags+=("${image_name}:${version}")
    
    # 如果是 latest 版本，同时添加不带版本的标签
    if [ "${version}" != "latest" ]; then
        image_tags+=("${image_name}:latest")
    fi
    
    print_info "构建 ${service} Docker 镜像..."
    print_info "服务目录：${service_dir}"
    print_info "镜像标签：${image_tags[*]}"
    
    # 构建命令参数
    local build_args=("-t" "${image_tags[0]}")
    
    # 添加更多标签
    for ((i=1; i<${#image_tags[@]}; i++)); do
        build_args+=("-t" "${image_tags[$i]}")
    done
    
    # 是否不使用缓存
    if [ "${no_cache}" = true ]; then
        build_args+=("--no-cache")
    fi
    
    # 添加构建上下文
    build_args+=("${service_dir}")
    
    # 执行构建
    print_info "执行：docker build ${build_args[*]}"
    
    if docker build "${build_args[@]}"; then
        print_info "${service} Docker 镜像构建成功！"
        
        # 显示镜像信息
        echo ""
        print_info "镜像信息:"
        docker images "${image_name}" | head -n 2
        echo ""
        
        return 0
    else
        print_error "${service} Docker 镜像构建失败！"
        return 1
    fi
}

# 主函数
main() {
    local service=""
    local version="${DEFAULT_VERSION}"
    local push=false
    local no_cache=false
    
    # 解析命令行参数
    while [[ $# -gt 0 ]]; do
        case $1 in
            -s|--service)
                service="$2"
                shift 2
                ;;
            -v|--version)
                version="$2"
                shift 2
                ;;
            -p|--push)
                push=true
                shift
                ;;
            -n|--no-cache)
                no_cache=true
                shift
                ;;
            -h|--help)
                show_help
                exit 0
                ;;
            *)
                print_error "未知选项：$1"
                show_help
                exit 1
                ;;
        esac
    done
    
    # 检查 Docker 是否可用
    if ! command -v docker &> /dev/null; then
        print_error "Docker 未安装或未在 PATH 中"
        exit 1
    fi
    
    # 检查 Docker 是否运行
    if ! docker info &> /dev/null; then
        print_error "Docker 未运行或无法连接"
        exit 1
    fi
    
    print_info "项目根目录：${PROJECT_ROOT}"
    print_info "构建版本：${version}"
    echo ""
    
    # 构建指定服务或所有服务
    if [ -n "${service}" ]; then
        # 构建单个服务
        if build_local "${service}" && build_service "${service}" "${version}" "${no_cache}"; then
            print_info "构建完成！"
        else
            print_error "构建失败！"
            exit 1
        fi
    else
        # 构建所有服务
        local failed=0
        for svc in ${SERVICES}; do
            if ! build_local "${svc}" || ! build_service "${svc}" "${version}" "${no_cache}"; then
                ((failed++))
            fi
        done
        
        echo ""
        if [ ${failed} -eq 0 ]; then
            print_info "所有服务构建完成！"
        else
            print_error "有 ${failed} 个服务构建失败"
            exit 1
        fi
    fi
    
    # 显示所有构建的镜像
    echo ""
    print_info "已构建的镜像:"
    docker images | grep "micro-go"
}

# 执行主函数
main "$@"
