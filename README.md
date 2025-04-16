# WASM-Go 沙箱示例

这个项目展示了如何使用 Go 和 WebAssembly 创建一个安全的沙箱环境，通过 wazero 运行时来执行 WASM 模块。

## 项目结构

```
.
├── math.go          # WASM 模块源代码
├── math.wasm        # 编译后的 WASM 模块
├── sandbox.go       # Go 沙箱运行时
└── README.md        # 项目文档
```

## 环境要求

- Go 1.23 或更高版本
- TinyGo 0.37.0 或更高版本
- wazero 运行时库

## 安装 TinyGo

### Ubuntu/Debian

```bash
wget https://github.com/tinygo-org/tinygo/releases/download/v0.37.0/tinygo_0.37.0_amd64.deb
sudo dpkg -i tinygo_0.37.0_amd64.deb
```

### 验证安装

```bash
tinygo version
```

## 编译 WASM 模块

使用以下命令将 Go 源代码编译为 WASM 模块：

```bash
tinygo build -o math.wasm -target=wasi math.go
```

## 运行沙箱

执行以下命令运行沙箱环境：

```bash
go run sandbox.go
```

## 项目特性

- 安全的沙箱环境
- 内存限制（~18MB）
- 执行超时保护
- WASI 支持
- 模块命名空间隔离

## WASM 模块导出函数

当前 WASM 模块导出了以下函数：

- `add(a, b int32) int32`: 计算两个整数的和

## 版本信息

- Go: 1.23+
- TinyGo: 0.37.0
- wazero: 最新版本

## TinyGo 相关说明

### 支持的目标平台

TinyGo 支持多个编译目标：

- wasi: WebAssembly 系统接口
- wasm: 浏览器 WebAssembly
- 其他嵌入式平台

### 编译参数说明

- `-target`: 指定目标平台
- `-o`: 指定输出文件
- `-gc`: 指定垃圾回收器类型
- `-scheduler`: 指定调度器类型
- `-size`: 显示二进制大小信息

### 导出函数注解

在 Go 代码中使用 `//export` 注解来导出函数到 WASM：

```go
//export add
func add(a, b int32) int32 {
    return a + b
}
```

## 注意事项

1. 确保 WASM 模块使用正确的导出注解
2. 注意内存限制和执行超时设置
3. 正确处理错误和资源清理

## 参考链接

- [TinyGo 官方文档](https://tinygo.org/)
- [wazero 文档](https://wazero.io/)
- [WebAssembly 系统接口](https://wasi.dev/) 