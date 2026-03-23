# dashboard-rest 目录结构优化对比

## 优化前目录结构

```
dashboard-rest/
├── api/
│   ├── dashboard_handler.go
│   └── router.go
├── build/
│   └── micro-go
├── cmd/
│   └── server/
│       └── main.go
├── grpc/
│   ├── client/
│   │   └── dashboard_client.go
│   └── proto/
│       ├── dashboard.pb.go
│       ├── dashboard.proto
│       └── dashboard_grpc.pb.go
├── internal/
│   ├── api/
│   │   ├── dashboard_handler.go
│   │   └── router.go
│   ├── grpc/
│   │   ├── client/
│   │   │   └── dashboard_client.go
│   │   └── proto/
│   │       ├── dashboard.pb.go
│   │       ├── dashboard.proto
│   │       └── dashboard_grpc.pb.go
│   └── model/
│       └── dashboard.go
├── model/
│   └── dashboard.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   └── logger/
│       └── logger.go
├── Makefile
├── go.mod
├── go.sum
└── main.go
```

## 优化后目录结构

```
dashboard-rest/
├── api/
│   ├── dashboard_handler.go
│   └── router.go
├── grpc/
│   ├── client/
│   │   └── dashboard_client.go
│   └── proto/
│       ├── dashboard.pb.go
│       ├── dashboard.proto
│       └── dashboard_grpc.pb.go
├── model/
│   └── dashboard.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   └── logger/
│       └── logger.go
├── validator/
├── configs/
├── scripts/
├── Makefile
├── go.mod
├── go.sum
└── main.go
```

## 优化内容说明

1. **移除的目录**：
   - `internal/`：移除了重复的内部目录，避免代码重复
   - `build/`：移除了构建输出目录，由 Makefile 动态生成
   - `cmd/server/`：移除了多余的命令目录，直接使用根目录的 main.go 作为入口

2. **添加的目录**：
   - `validator/`：添加了参数校验目录，用于请求参数校验
   - `configs/`：添加了配置文件目录，用于存放配置文件
   - `scripts/`：添加了脚本目录，用于存放各种脚本

3. **修改的文件**：
   - `Makefile`：更新了构建脚本，使用 bin/ 目录而不是 build/ 目录
   - `go.mod`：添加了 Gin、Validator 和 Redis 依赖项
   - `model/dashboard.go`：统一了字段命名，将 "开户人数" 改为 "accountCount"
   - `api/router.go`：使用 Gin 框架来设置路由
   - `api/dashboard_handler.go`：将 HTTP 处理函数改为使用 Gin 框架的上下文
   - `main.go`：更新了路由设置，使用 Gin 框架的 Run 方法启动服务器

4. **保持不变的目录**：
   - `api/`：API 层目录，处理 HTTP 请求和响应
   - `grpc/`：gRPC 客户端目录，用于调用业务/桥接服务
   - `model/`：数据模型目录，定义数据结构
   - `pkg/`：公共包目录，包含配置、日志等公共组件

## 优化效果

1. **目录结构更加清晰**：移除了重复和多余的目录，添加了必要的目录，使目录结构更加符合项目规范
2. **代码组织结构更加合理**：使用 Gin 框架替代了标准的 HTTP 处理，提高了代码的可维护性和可扩展性
3. **技术栈更加符合规范**：添加了 Gin、Validator 和 Redis 依赖项，确保技术栈符合项目要求
4. **构建脚本更加规范**：更新了 Makefile 文件，使其符合项目规范要求
5. **配置文件和依赖项更加清晰**：确保所有配置文件和依赖项都符合项目规范

通过这些优化，dashboard-rest 目录现在更加符合项目架构规范，代码更加可维护和可扩展，为后续的开发和部署做好了准备。