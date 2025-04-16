# 常用命令指南

## TinyGo 相关命令

### 编译 WASM 文件
```bash
# 编译为 WASI 目标
tinygo build -o math.wasm -target=wasi math.go

# 编译为浏览器 WASM 目标
tinygo build -o math.wasm -target=wasm math.go

# 显示编译大小信息
tinygo build -size -o math.wasm -target=wasi math.go

# 优化编译（减小文件大小）
tinygo build -opt=2 -o math.wasm -target=wasi math.go
```

### 查看版本信息
```bash
# 查看 TinyGo 版本
tinygo version

# 查看支持的目标平台列表
tinygo targets
```

## Go 相关命令

### 项目初始化和依赖管理
```bash
# 初始化 Go 模块
go mod init wasm-go

# 添加/更新依赖
go get github.com/tetratelabs/wazero
go get github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1

# 整理依赖
go mod tidy
```

### 运行和测试
```bash
# 运行沙箱程序
go run sandbox.go

# 运行测试
go test ./...

# 带详细输出的测试
go test -v ./...
```

## WASM 调试命令

### 查看 WASM 文件信息
```bash
# 查看导出函数列表（需要安装 wasm-tools）
wasm-tools print math.wasm

# 查看文件大小
ls -lh math.wasm
```

### 性能分析
```bash
# 运行时性能分析
go run -cpuprofile cpu.prof sandbox.go

# 内存分析
go run -memprofile mem.prof sandbox.go
```

## 常用开发工作流

### 完整的构建和运行流程
```bash
# 1. 清理旧的构建文件
rm -f math.wasm

# 2. 编译 WASM 模块
tinygo build -o math.wasm -target=wasi math.go

# 3. 运行沙箱程序
go run sandbox.go
```

### 开发调试流程
```bash
# 1. 启用详细日志
export WAZEROLOG=debug

# 2. 运行程序
go run sandbox.go

# 3. 关闭详细日志
unset WAZEROLOG
```

## 环境设置

### TinyGo 环境变量
```bash
# 设置 TINYGOROOT（如果需要）
export TINYGOROOT=/usr/local/lib/tinygo

# 添加 TinyGo 到 PATH（如果需要）
export PATH=$PATH:/usr/local/lib/tinygo/bin
```

### 代理设置（如果需要）
```bash
# 设置 Go 代理
export GOPROXY=https://goproxy.cn,direct

# 设置私有模块（如果使用私有仓库）
export GOPRIVATE=*.internal.example.com
```

## 故障排除命令

### 清理和重置
```bash
# 清理构建缓存
go clean -cache

# 清理模块缓存
go clean -modcache

# 重新下载所有依赖
rm -rf go.sum && go mod download
```

### 诊断工具
```bash
# 检查 Go 环境
go env

# 检查 TinyGo 环境
tinygo env

# 验证模块依赖
go mod verify
```