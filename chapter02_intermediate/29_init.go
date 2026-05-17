// 本课目标：了解 init 函数与包初始化顺序。
//
// 知识点：
//   - 每个包可有多个 init，按源文件顺序执行
//   - 顺序：import 的包 init → 当前包变量初始化 → 当前包 init
//   - main 包中 init 在 main 之前执行
//   - init 不能手动调用
//
// 运行：go run ./chapter02_intermediate/29_init.go
package main

import "fmt"

var order = appendOrder("包级变量初始化")

func appendOrder(step string) string {
	fmt.Println(step)
	return step
}

func init() {
	fmt.Println("init 函数执行")
}

func main() {
	fmt.Println("main 函数执行")
}
