前端服务，开发过程中，禁止使用模拟数据

rest 服务，禁止使用 mock 数据，data 层可以使用 mock数据

通过proto 文件生成 代码时，参考指令   protoc --go\_out=../../ --go-grpc\_out=../../ dashboard.proto

严格 按 ARCHITECTURE.md文档要求编写代码

docker 构建时，禁止采用多阶段构建。所有服务在本地编译，再打包到 docker 镜像
后服务编译时，应有必要的 log 输出，方便调试。