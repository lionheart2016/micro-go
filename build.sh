#!/bin/bash

# 编译脚本，为所有 Golang 服务生成 amd64 架构的可执行文件

echo "开始编译所有 Golang 服务..."

# 设置编译环境变量，生成 amd64 架构的可执行文件
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

# 编译 dashboard-data 服务
echo "编译 dashboard-data 服务..."
cd dashboard-data
mkdir -p bin
rm -f bin/server
go build -o bin/server main.go
if [ $? -eq 0 ]; then
    echo "dashboard-data 服务编译成功"
else
    echo "dashboard-data 服务编译失败"
    exit 1
fi
cd ..

# 编译 dashboard-rest 服务
echo "编译 dashboard-rest 服务..."
cd dashboard-rest
mkdir -p bin
rm -f bin/server
go build -o bin/server main.go
if [ $? -eq 0 ]; then
    echo "dashboard-rest 服务编译成功"
else
    echo "dashboard-rest 服务编译失败"
    exit 1
fi
cd ..

# 编译 sample-rest 服务
echo "编译 sample-rest 服务..."
cd sample-rest
mkdir -p bin
rm -f bin/server
go build -o bin/server main.go
if [ $? -eq 0 ]; then
    echo "sample-rest 服务编译成功"
else
    echo "sample-rest 服务编译失败"
    exit 1
fi
cd ..

echo "所有 Golang 服务编译完成！"
