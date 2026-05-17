// 本课目标：掌握类型断言与 type switch。
//
// 知识点：
//   - v, ok := x.(T)：断言 x 的动态类型是否为 T
//   - switch x.(type) { case T: ... } 多类型分支
//   - 断言失败且不用 ok 时会 panic
//
// 运行：go run ./chapter02_intermediate/23_type_assertion.go
package main

import "fmt"

func describe(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("整数", v)
	case string:
		fmt.Println("字符串", v)
	case bool:
		fmt.Println("布尔", v)
	default:
		fmt.Printf("未知类型 %T\n", v)
	}
}

func main() {
	var x any = "Go"

	// 单类型断言
	s, ok := x.(string)
	if ok {
		fmt.Println("是 string:", s)
	}

	// type switch
	describe(42)
	describe("hello")
	describe(true)
	describe(3.14)
}
