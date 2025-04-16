package main

import (
	"fmt"
	"unsafe"
)

//export add
func add(a, b int32) int32 {
	return a + b
}

//export log
func log(msgPtr int32) {
	fmt.Println("[WASM] 开始读取内存...")
	fmt.Printf("[WASM] 内存指针: %d\n", msgPtr)

	// 从 WASM 内存读取字符串
	msg := readString(msgPtr)
	fmt.Printf("[WASM] 读取到的消息: %s\n", msg)
	fmt.Println("[WASM] 日志输出完成")
}

// 从内存指针读取字符串
func readString(ptr int32) string {
	// 获取字符串长度（假设字符串以 null 结尾）
	var length int
	p := unsafe.Pointer(uintptr(ptr))

	fmt.Printf("[WASM] 开始查找字符串长度，起始地址: %d\n", ptr)
	for {
		if *(*byte)(unsafe.Pointer(uintptr(p) + uintptr(length))) == 0 {
			break
		}
		length++
		if length > 1000 { // 安全检查
			fmt.Println("[WASM] 警告：字符串过长，可能存在问题")
			break
		}
	}
	fmt.Printf("[WASM] 字符串长度: %d\n", length)

	// 转换为字符串
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = *(*byte)(unsafe.Pointer(uintptr(p) + uintptr(i)))
	}
	return string(bytes)
}

func main() {} // 必须保留空 main 函数
