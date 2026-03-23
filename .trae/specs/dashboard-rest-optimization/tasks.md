# dashboard-rest 目录优化 - 实施计划

## [x] 任务 1: 分析当前目录结构
- **Priority**: P0
- **Depends On**: None
- **Description**:
  - 分析当前 dashboard-rest 目录结构
  - 与项目规范要求进行对比
  - 识别需要调整的部分
- **Acceptance Criteria Addressed**: AC-1
- **Test Requirements**:
  - `human-judgment` TR-1.1: 确认当前目录结构与规范的差异
- **Notes**: 需要特别关注重复的 internal/ 目录和缺少的目录
- **完成情况**:
  - 当前目录结构存在重复的 internal/ 目录
  - 缺少 validator/、configs/、scripts/ 目录
  - 存在不需要的 build/ 和 cmd/ 目录
  - 其他目录基本符合规范

## [x] 任务 2: 调整目录结构
- **Priority**: P0
- **Depends On**: 任务 1
- **Description**:
  - 移除重复的 internal/ 目录
  - 确保目录结构符合项目规范
  - 添加缺少的目录（如 validator/、configs/、scripts/）
- **Acceptance Criteria Addressed**: AC-1
- **Test Requirements**:
  - `human-judgment` TR-2.1: 确认目录结构符合项目规范
- **Notes**: 注意保留必要的文件，避免删除重要代码
- **完成情况**:
  - 移除了 internal/、build/、cmd/ 目录
  - 添加了 validator/、configs/、scripts/ 目录
  - 目录结构现在符合项目规范要求

## [x] 任务 3: 检查并更新技术栈
- **Priority**: P1
- **Depends On**: 任务 2
- **Description**:
  - 检查当前使用的技术栈
  - 确保使用 Golang + Gin、Apifox、Validator 和 Redis
  - 更新必要的依赖项
- **Acceptance Criteria Addressed**: AC-2
- **Test Requirements**:
  - `human-judgment` TR-3.1: 确认技术栈符合项目规范
- **Notes**: 检查 go.mod 文件中的依赖项
- **完成情况**:
  - 更新了 go.mod 文件，添加了 Gin、Validator 和 Redis 依赖项
  - 运行了 go mod tidy 命令更新 go.sum 文件
  - 技术栈现在符合项目规范要求

## [x] 任务 4: 优化配置文件和依赖项
- **Priority**: P1
- **Depends On**: 任务 3
- **Description**:
  - 检查并更新配置文件
  - 确保依赖项符合项目规范
  - 更新构建脚本
- **Acceptance Criteria Addressed**: AC-3
- **Test Requirements**:
  - `human-judgment` TR-4.1: 确认配置文件和依赖项符合项目规范
- **Notes**: 检查 Makefile 和其他配置文件
- **完成情况**:
  - 更新了 Makefile 文件，使其符合项目规范
  - 检查了 main.go、api/dashboard_handler.go、api/router.go、grpc/client/dashboard_client.go、model/dashboard.go、pkg/config/config.go 和 pkg/logger/logger.go 文件
  - 所有配置文件和依赖项现在符合项目规范要求

## [x] 任务 5: 优化代码组织结构
- **Priority**: P1
- **Depends On**: 任务 4
- **Description**:
  - 优化代码组织结构
  - 确保代码符合命名规范和接口规范
  - 提高代码可维护性和可扩展性
- **Acceptance Criteria Addressed**: AC-4
- **Test Requirements**:
  - `human-judgment` TR-5.1: 确认代码组织结构合理
- **Notes**: 注意保持代码的业务逻辑不变
- **完成情况**:
  - 修改了 model/dashboard.go 文件，统一了字段命名，将 "开户人数" 改为 "accountCount"
  - 修改了 api/router.go 文件，使用 Gin 框架来设置路由
  - 修改了 api/dashboard_handler.go 文件，将 HTTP 处理函数改为使用 Gin 框架的上下文
  - 修改了 main.go 文件，更新了路由设置，使用 Gin 框架的 Run 方法启动服务器
  - 代码组织结构现在更加合理，符合项目规范要求

## [x] 任务 6: 生成目录结构对比说明
- **Priority**: P2
- **Depends On**: 任务 5
- **Description**:
  - 生成优化前后的目录结构对比说明
  - 清晰展示优化的内容和效果
- **Acceptance Criteria Addressed**: AC-5
- **Test Requirements**:
  - `human-judgment` TR-6.1: 确认目录结构对比说明清晰明了
- **Notes**: 可以使用文本或表格形式展示对比
- **完成情况**:
  - 创建了 directory-structure-comparison.md 文件，详细记录了优化前后的目录结构
  - 清晰展示了优化的内容和效果，包括移除的目录、添加的目录、修改的文件和保持不变的目录
  - 说明了优化后的效果，包括目录结构更加清晰、代码组织结构更加合理、技术栈更加符合规范等

## [x] 任务 7: 构建和运行服务
- **Priority**: P0
- **Depends On**: 任务 6
- **Description**:
  - 构建优化后的 dashboard-rest 服务
  - 运行服务并验证其正常工作
- **Acceptance Criteria Addressed**: AC-6
- **Test Requirements**:
  - `programmatic` TR-7.1: 确认服务能够正常构建
  - `programmatic` TR-7.2: 确认服务能够正常运行
- **Notes**: 检查服务启动是否成功，是否有错误信息
- **完成情况**:
  - 成功构建了 dashboard-rest 服务
  - 成功启动了服务，监听在端口 8082 上
  - Gin 框架已经成功初始化，并且注册了两个路由：GET /api/core-indicators 和 GET /api/performance
  - 服务运行正常，没有错误信息