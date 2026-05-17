// 本课目标：掌握 if/else 及「带初始化语句的 if」。
//
// 知识点：
//   - 条件不需要括号，但花括号必须写
//   - if 可带 init 语句：if v := expr(); v > 0 { }
//   - else if / else 与 if 写在同一列（go fmt 会格式化）
//
// 运行：go run ./chapter01_basics/07_if_else.go
package main

import "fmt"

func main() {
	score := 85

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// 带初始化的 if：变量作用域仅在 if-else 块内
	if n := score / 10; n >= 8 {
		fmt.Println("十分位 >= 8, n =", n)
	} else {
		fmt.Println("十分位 < 8, n =", n)
	}
	// fmt.Println(n)  // 编译错误：undefined: n

	// 常见模式：先调用返回 (value, err)，在 if 里判断 err
	value, ok := divide(10, 2)
	if !ok {
		fmt.Println("除法失败")
	} else {
		fmt.Println("10/2 =", value)
	}

	_, ok = divide(10, 0)
	if !ok {
		fmt.Println("除零被拦截")
	}
}

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}
