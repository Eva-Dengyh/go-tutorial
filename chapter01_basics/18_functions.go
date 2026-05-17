// 本课目标：掌握函数定义、多返回值、命名返回值与函数类型。
//
// 知识点：
//   - func 名(参数) 返回值 { }
//   - 多返回值常用于 (结果, error)
//   - 命名返回值在 return 时可省略变量名（慎用）
//   - 函数是一等公民：可赋值给变量、作为参数传递
//
// 运行：go run ./chapter01_basics/18_functions.go
package main

import "fmt"

// 普通函数
func add(a, b int) int {
	return a + b
}

// 多返回值
func divMod(a, b int) (int, int) {
	return a / b, a % b
}

// 命名返回值
func swap(a, b int) (x, y int) {
	x, y = b, a
	return // 裸 return，返回 x, y
}

func main() {
	fmt.Println("add(3,5) =", add(3, 5))

	q, r := divMod(10, 3)
	fmt.Printf("10/3 商=%d 余=%d\n", q, r)

	x, y := swap(1, 2)
	fmt.Println("swap 后", x, y)

	// 函数类型
	var op func(int, int) int = add
	fmt.Println("op(2,3) =", op(2, 3))

	// 匿名函数
	double := func(n int) int { return n * 2 }
	fmt.Println("double(5) =", double(5))
}
