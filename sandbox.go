// sandbox.go
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	ctx := context.Background()

	fmt.Println("开始读取 WASM 文件...")
	// 读取 WASM 文件
	math_wasm, err := os.ReadFile("math.wasm")
	if err != nil {
		panic(fmt.Sprintf("读取文件失败: %v", err))
	}
	fmt.Printf("WASM 文件读取成功，大小: %d 字节\n", len(math_wasm))

	// 创建运行时（开启沙箱模式）
	fmt.Println("创建 WASM 运行时...")
	runtime := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfig().
		WithMemoryLimitPages(300).    // 限制最大内存为 ~18MB (300 * 64KB)
		WithCloseOnContextDone(true)) // 严格上下文隔离

	// 初始化 WASI 子系统
	fmt.Println("初始化 WASI...")
	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

	// 加载 WASM 模块（带安全策略）
	fmt.Println("加载 WASM 模块...")
	config := wazero.NewModuleConfig().
		WithName("math").    // 模块命名空间隔离
		WithStartFunctions() // 禁止自动执行_start

	mod, err := runtime.InstantiateWithConfig(ctx, math_wasm, config)
	if err != nil {
		panic(fmt.Sprintf("加载模块失败: %v", err))
	}
	fmt.Println("WASM 模块加载成功")

	// 设置执行超时（防止无限循环）
	ctx, cancel := context.WithTimeout(ctx, 5000*time.Second)
	defer cancel()

	// 执行导出函数（带参数验证）
	fmt.Println("\n=== 测试 add 函数 ===")
	if addFunc := mod.ExportedFunction("add"); addFunc != nil {
		fmt.Println("开始执行 add 函数...")
		result, err := addFunc.Call(ctx, 10, 20)
		if err != nil {
			fmt.Println("执行失败:", err)
		} else {
			fmt.Println("计算结果:", result[0])
		}
	} else {
		fmt.Println("未找到 add 函数")
	}

	// 测试 log 函数
	fmt.Println("\n=== 测试 log 函数 ===")
	if logFunc := mod.ExportedFunction("log"); logFunc != nil {
		fmt.Println("开始执行 log 函数...")

		// 获取模块内存
		memory := mod.Memory()
		if memory == nil {
			panic("无法访问 WASM 模块内存")
		}

		// 写入测试字符串到 WASM 内存
		testMsg := "你好，这是来自宿主环境的消息！\x00" // 添加 null 终止符
		msgBytes := []byte(testMsg)

		// 写入字符串数据到内存起始位置
		if !memory.Write(0, msgBytes) {
			panic("写入内存失败")
		}

		// 调用 log 函数，传递字符串起始地址
		_, err = logFunc.Call(ctx, 0)
		if err != nil {
			fmt.Println("执行失败:", err)
		}
		fmt.Println("执行成功log函数")
	} else {
		fmt.Println("未找到 log 函数")
	}

	// 清理沙箱
	runtime.Close(ctx)
}
