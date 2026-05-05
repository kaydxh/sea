<p align="center">
  <h1 align="center">🌊 Sea</h1>
  <p align="center">A production-ready Go microservice scaffold based on gRPC-Gateway and Clean Architecture</p>
</p>

<p align="center">
  <a href="https://github.com/kaydxh/sea/actions"><img src="https://img.shields.io/badge/build-passing-brightgreen.svg?style=flat-square" alt="Build Status"></a>
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.25+-00ADD8.svg?style=flat-square&logo=go" alt="Go Version"></a>
  <a href="https://grpc.io"><img src="https://img.shields.io/badge/gRPC-Gateway-blue.svg?style=flat-square" alt="gRPC-Gateway"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License"></a>
</p>

---

## 简介

**Sea** 是一个基于 Go 语言的微服务脚手架框架，集成了 gRPC + gRPC-Gateway（同时暴露 gRPC 和 RESTful HTTP 接口）、Clean Architecture 分层设计、OpenTelemetry 可观测性、以及开箱即用的中间件体系。

Sea 的设计目标是：**一键创建新项目，开箱即用，专注业务逻辑开发**。

## ✨ 特性

- 🚀 **gRPC + HTTP 双协议** — 基于 gRPC-Gateway，一份 Proto 定义同时生成 gRPC 和 RESTful API
- 🏗️ **Clean Architecture** — 严格分层（Domain → Application → Infrastructure → Web），依赖反转，易于测试
- 📡 **OpenTelemetry** — 内置 Trace、Metrics、Logging 全链路可观测性（支持 Jaeger / Prometheus / OTLP）
- 🔌 **插件化配置** — MySQL、Redis、日志、服务发现等组件通过 Plugin 模式按需加载
- 🛡️ **生产级中间件** — CORS、请求限流（QPS/并发）、Debug Body 打印、健康检查
- 🎯 **一键脚手架** — `make new TARGET=your-service` 即可创建完整的新微服务项目
- 📦 **Docker 支持** — 内置 Dockerfile 和部署脚本
- ⚙️ **环境变量覆盖** — 敏感配置支持环境变量注入，适配容器化部署

## 📐 架构设计

```
┌─────────────────────────────────────────────────────────────────┐
│                         cmd/ (入口层)                            │
│  main.go → server.go → options/ (Plugin 插件注册)                │
└────────────────────────────────┬────────────────────────────────┘
                                 │
┌────────────────────────────────▼────────────────────────────────┐
│                         web/ (接口层)                            │
│  Router (中间件注册) → Controller (RPC 方法实现) → Error 映射     │
└────────────────────────────────┬────────────────────────────────┘
                                 │
┌────────────────────────────────▼────────────────────────────────┐
│                    pkg/application/ (应用层)                      │
│  Handler (编排业务用例，持有 Repository 接口)                     │
└────────────────────────────────┬────────────────────────────────┘
                                 │
┌────────────────────────────────▼────────────────────────────────┐
│                     pkg/domain/ (领域层)                          │
│  Entity + Repository 接口 + Domain Error (纯业务逻辑，无外部依赖) │
└────────────────────────────────┬────────────────────────────────┘
                                 │
┌────────────────────────────────▼────────────────────────────────┐
│                 pkg/infrastructure/ (基础设施层)                   │
│  MySQL DAO / Redis DAO / Local 实现 (实现 Repository 接口)        │
└─────────────────────────────────────────────────────────────────┘
```

## 📁 项目结构

```
sea/
├── api/                          # Proto 接口定义及生成代码
│   └── protoapi-spec/
│       └── sea-date/v1/
│           ├── api.proto         # 服务接口定义（含 HTTP 路由注解）
│           ├── configuration.proto # 配置结构定义
│           └── error.proto       # 业务错误码定义
├── cmd/                          # 应用入口
│   └── sea-date/
│       ├── main.go              # 程序入口
│       └── app/
│           ├── server.go        # 服务启动逻辑
│           └── options/         # 插件式配置加载
│               ├── plugin.config.go
│               ├── plugin.mysql.go
│               ├── plugin.redis.go
│               └── plugin.web_handler.go
├── conf/                         # 配置文件
│   └── sea-date.yaml
├── pkg/                          # 核心业务逻辑
│   └── sea-date/
│       ├── application/         # 应用层（用例编排）
│       ├── domain/              # 领域层（实体 + 接口）
│       ├── infrastructure/      # 基础设施层（接口实现）
│       │   ├── database/        # MySQL/Redis DAO
│       │   └── local/           # 本地实现
│       └── provider/            # 依赖注入容器
├── web/                          # Web 接口层
│   ├── app/                     # 路由注册入口
│   └── modules/                 # 业务模块（Controller + Router）
├── script/                       # 构建和工具脚本
├── test/                         # 集成测试
└── docs/                         # 项目文档
```

## 🚀 快速开始

### 环境要求

- Go 1.25+
- protoc (Protocol Buffers 编译器)
- protoc-gen-go, protoc-gen-go-grpc, protoc-gen-grpc-gateway

### 构建运行

```bash
# 克隆项目
git clone https://github.com/kaydxh/sea.git
cd sea

# 构建
make

# 运行
./cmd/sea-date/sea-date --config=./conf/sea-date.yaml
```

### 生成 Proto 代码

```bash
make generate
```

### 测试接口

```bash
# HTTP (gRPC-Gateway)
curl -X POST http://localhost:10001/api/now \
  -H "Content-Type: application/json" \
  -d '{"RequestId": "test-001"}'

# gRPC (使用 grpcurl)
grpcurl -plaintext -d '{"request_id": "test-001"}' \
  localhost:10001 sea.api.seadate.SeaDateService/Now
```

## 🎯 创建新项目

Sea 最强大的能力是作为脚手架快速创建新的微服务项目：

```bash
# 在 sea 项目根目录下执行
make new TARGET=your-service NEW_GIT_REPOSITORY_NAME="github.com/yourname/your-repo"
```

该命令会：
1. 下载 Sea 模板
2. 将 `sea-date` 重命名为 `sea-your-service`
3. 替换所有包路径为你的 Git 仓库地址
4. 生成完整的项目结构，开箱即用

## ⚙️ 配置说明

配置文件位于 `conf/sea-date.yaml`，支持以下模块：

| 模块 | 说明 |
|------|------|
| `log` | 日志配置（级别、轮转、输出路径） |
| `web` | HTTP/gRPC 服务配置（端口、超时、API 格式化） |
| `web.open_telemetry` | OpenTelemetry 可观测性（Trace/Metrics/Logs） |
| `web.qps_limit` | QPS 限流配置（gRPC/HTTP 分别配置） |
| `resolver` | 服务发现配置（DNS/Consul） |
| `database.mysql` | MySQL 连接配置 |
| `database.redis` | Redis 连接配置 |
| `debug` | 调试配置（dump 等） |

### 环境变量覆盖

敏感配置支持通过环境变量覆盖，适配 Kubernetes / Docker 部署：

```bash
export DB_ADDRESS="mysql-host:3306"
export DB_USERNAME="admin"
export DB_PASSWORD="secret"
export REDIS_ADDRESS="redis-host:6379"
export REDIS_PASSWORD="secret"
```

## 🔧 常用命令

```bash
# 构建项目
make

# 生成 Proto 代码 + 文档
make generate

# 创建新服务
make new TARGET=<service-name> NEW_GIT_REPOSITORY_NAME="<git-path>"

# 删除服务
make delete TARGET=<service-name>

# 打包部署
make install

# 清理构建产物
make clean
```

## 🧩 核心依赖

| 依赖 | 用途 |
|------|------|
| [kaydxh/golang](https://github.com/kaydxh/golang) | 基础库（日志、Web 框架、中间件、数据库封装） |
| [grpc-gateway/v2](https://github.com/grpc-ecosystem/grpc-gateway) | gRPC-Gateway HTTP 代理 |
| [spf13/cobra](https://github.com/spf13/cobra) | CLI 命令行框架 |
| [OpenTelemetry](https://opentelemetry.io/) | 分布式追踪和指标采集 |
| [jmoiron/sqlx](https://github.com/jmoiron/sqlx) | MySQL 数据库操作 |
| [go-redis/redis](https://github.com/go-redis/redis) | Redis 客户端 |

## 📖 文档

- [API 接口文档](docs/api.md)
- [配置文档](docs/configuration.md)
- [错误码文档](docs/error_code.md)
- [业务模块文档](docs/date.md)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'feat: add amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 📄 License

本项目采用 [MIT License](LICENSE) 开源协议。