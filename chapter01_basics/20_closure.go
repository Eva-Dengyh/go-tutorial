// 本课目标：理解闭包如何捕获外部变量，以及常见陷阱。
//
// 知识点：
//   - 闭包：函数引用其定义作用域外的变量
//   - 返回的函数仍持有对外层变量的引用
//   - 循环中创建闭包要注意：应传参复制 i，避免都引用同一变量
//
// 运行：go run ./chapter01_basics/20_closure.go
package main

import "fmt"

// counter 返回一个每次调用自增 1 的函数（闭包持有 n）
func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	next := counter()
	fmt.Println(next(), next(), next()) // 1 2 3

	// 两个闭包各自独立的 n
	a, b := counter(), counter()
	fmt.Println("a:", a(), a(), "b:", b())

	// 陷阱：循环变量被多个闭包共享
	fmt.Print("错误示范: ")
	var funcs []func()
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() { fmt.Print(i, " ") })
	}
	for _, f := range funcs {
		f() // 可能都打印 3
	}
	fmt.Println()

	// 正确：用参数复制当前 i
	fmt.Print("正确示范: ")
	funcs = nil
	for i := 0; i < 3; i++ {
		i := i // 或作为参数传入立即执行的函数
		funcs = append(funcs, func() { fmt.Print(i, " ") })
	}
	for _, f := range funcs {
		f()
	}
	fmt.Println()
}
