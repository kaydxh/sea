# 问题总结

---

## Issue #1: Go Workspace 多模块仓库使用问题

### 问题背景

在使用 `go.work` 让 `sea` 项目依赖本地 `golang` 仓库代码时，遇到了一系列问题。

### 仓库结构

```
github.com/kaydxh/
├── go.work
├── golang/                    # 多模块仓库
│   ├── go.mod                 # 主模块: github.com/kaydxh/golang
│   ├── doc.go
│   ├── go/
│   │   └── go.mod             # 子模块: github.com/kaydxh/golang/go
│   └── pkg/
│       ├── config/
│       │   └── go.mod         # 子模块: github.com/kaydxh/golang/pkg/config
│       ├── webserver/
│       │   └── go.mod         # 子模块: github.com/kaydxh/golang/pkg/webserver
│       └── ...                # 其他子模块
└── sea/                       # 依赖 golang 的项目
    └── go.mod
```

### 遇到的问题

#### 1.1 go mod tidy 忽略 go.work

**现象**：执行 `make` 时仍然下载远程依赖

```bash
go: downloading github.com/kaydxh/golang v0.0.151
```

**原因**：`build.sh` 中执行了 `go mod tidy`，而 **`go mod tidy` 不受 go.work 影响**，会直接按 go.mod 中的版本下载远程依赖。

**解决方案**：修改 `build.sh`，检测到 go.work 存在时跳过 `go mod tidy`：

```bash
# 如果使用 go.work，跳过 go mod tidy（go mod tidy 会忽略 go.work）
if [[ -z "$(go env GOWORK)" ]]; then
  go mod tidy
fi
```

#### 1.2 unknown revision

**现象**：

```
github.com/kaydxh/golang@v0.0.151: unknown revision v0.0.151
```

**原因**：go.mod 中声明的版本号 `v0.0.151` 在远程仓库不存在（远程最新版本是 `v0.0.134`）。

**解决方案**：使用远程存在的版本号：

```go
require (
    github.com/kaydxh/golang v0.0.134
)
```

go.work 的 `use` 指令会自动用本地模块覆盖这个版本。

#### 1.3 use 和 replace 冲突

**现象**：

```
go: workspace module github.com/kaydxh/golang is replaced at all versions in the go.work file
```

**原因**：在 go.work 中同时使用 `use` 和 `replace` 指向同一个模块。

**解决方案**：只使用 `use`，不要添加 `replace`：

```go
// go.work
go 1.25.3

use (
    ./golang
    ./golang/go
    ./golang/pkg/config
    // ... 其他子模块
    ./sea
)
```

#### 1.4 子模块版本 v0.0.0 不存在

**现象**：

```
github.com/kaydxh/golang/pkg/config@v0.0.0: unknown revision pkg/config/v0.0.0
```

**原因**：尝试在 sea 的 go.mod 中直接依赖子模块并使用 `v0.0.0` 版本，但这个版本在远程不存在。

**解决方案**：不要在 go.mod 中直接依赖子模块，而是依赖主模块。go.work 会自动处理子模块的解析。

### 正确配置

#### go.work 文件

```go
go 1.25.3

use (
    // golang 主模块
    ./golang
    ./golang/go

    // golang/pkg 子模块
    ./golang/pkg/binlog
    ./golang/pkg/config
    ./golang/pkg/container
    ./golang/pkg/crontab
    ./golang/pkg/database
    ./golang/pkg/discovery
    ./golang/pkg/file-cleanup
    ./golang/pkg/file-rotate
    ./golang/pkg/file-transfer
    ./golang/pkg/fsnotify
    ./golang/pkg/gocv
    ./golang/pkg/grpc-gateway
    ./golang/pkg/logs
    ./golang/pkg/middleware
    ./golang/pkg/monitor
    ./golang/pkg/mq
    ./golang/pkg/pool
    ./golang/pkg/profile
    ./golang/pkg/protobuf
    ./golang/pkg/proxy
    ./golang/pkg/resolver
    ./golang/pkg/scheduler
    ./golang/pkg/storage
    ./golang/pkg/viper
    ./golang/pkg/webserver

    // sea 项目
    ./sea
)
```

#### sea/go.mod 文件

```go
module github.com/kaydxh/sea

go 1.25.3

require (
    github.com/kaydxh/golang v0.0.134  // 使用远程存在的版本
    // ... 其他依赖
)
```

### 关键要点

1. **go.work 的 `use` 指令**：会自动将本地模块覆盖远程依赖，无需 `replace`
2. **go.mod 版本号必须有效**：即使使用 go.work，go.mod 中的版本号仍需在远程存在（用于验证）
3. **go mod tidy 不受 go.work 影响**：会按 go.mod 下载远程依赖，使用 go.work 开发时应跳过
4. **多模块仓库**：需要在 go.work 中列出所有子模块，go.work 会自动处理模块间的依赖关系

### 验证方法

```bash
# 检查 go.work 是否生效
go env GOWORK

# 直接编译测试
cd sea/cmd/sea-date && go build -o sea-date

# 查看模块解析结果
go list -m all | grep kaydxh
```

---
