---
name: 项目编译
description: 用于项目的编译、构建和部署，包括Go服务和前端服务的编译打包
---

# 项目编译技能

## 功能描述
- 编译Go服务（dashboard-data、dashboard-rest、sample-rest）
- 构建前端服务
- 构建Docker镜像
- 部署服务到本地环境

## 使用方法

### 1. 编译Go服务
```bash
./build.sh
```

### 2. 构建前端服务
```bash
cd frontend
yarn install
yarn build
```

### 3. 构建并部署服务
使用docker-compose构建并部署所有服务：
```bash
docker-compose up -d --build
```

或者使用已构建的镜像部署：
```bash
docker-compose up -d
```

## 编译环境
- Go 1.25+
- Node.js 18+
- Yarn
- Docker
- Docker Compose

## 注意事项
- 编译Go服务时会生成amd64架构的可执行文件，适用于Docker容器
- 前端服务使用Vite构建工具
- 所有服务通过Docker Compose部署在同一网络中
- 服务启动后可通过以下地址访问：
  - 前端：http://localhost:80
  - REST服务：http://localhost:8082
  - gRPC服务：localhost:50051