// 本课目标：理解 defer、panic、recover 的执行顺序与使用场景。
//
// 知识点：
//   - defer：函数返回前执行，多个 defer 后进先出（栈）
//   - panic：引发运行时恐慌，未 recover 则程序崩溃
//   - recover：仅在 defer 中有效，可捕获 panic
//   - 业务错误用 error，真正异常才考虑 panic
//
// 运行：go run ./chapter02_intermediate/27_defer_panic_recover.go
package main

import "fmt"

func deferDemo() {
	fmt.Println("进入 deferDemo")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("即将返回")
}

func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover 捕获:", r)
		}
	}()
	panic("出错了")
}

func main() {
	deferDemo()
	fmt.Println("---")
	mayPanic()
	fmt.Println("程序继续执行")
}
