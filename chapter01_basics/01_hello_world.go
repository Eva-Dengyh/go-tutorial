// 本课目标：理解 Go 程序的最小结构，能写出并运行第一个程序。
//
// 知识点：
//   - package：每个 Go 文件属于一个包；可执行程序入口包名必须是 main
//   - import：导入标准库或第三方包
//   - func main()：程序从这里开始执行（无参数、无返回值）
//   - fmt.Println：向标准输出打印并换行
//
// 运行：go run ./chapter01_basics/01_hello_world.go
package main

import "fmt"

func main() {
	// Println 可接收多个参数，以空格分隔，末尾自动换行
	fmt.Println("Hello, Go!")

	// Printf 按格式字符串输出；%s 字符串，%d 整数，%v 默认格式
	name := "学习者"
	year := 2026
	fmt.Printf("欢迎 %s，今年是 %d 年\n", name, year)

	// 包名后的点号：调用该包导出的函数（首字母大写才可被外部使用）
	fmt.Print("同一行")
	fmt.Print("不换行\n")
}
