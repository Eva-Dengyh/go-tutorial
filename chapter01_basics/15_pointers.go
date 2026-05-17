// 本课目标：理解指针 &、*，以及 new 与 make 的区别。
//
// 知识点：
//   - & 取地址，* 解引用
//   - 函数传指针可修改调用方变量
//   - new(T) 分配零值并返回 *T（很少单独用）
//   - make 只用于 slice、map、channel，返回已初始化可用的引用类型
//
// 运行：go run ./chapter01_basics/15_pointers.go
package main

import "fmt"

func addOne(n *int) {
	*n++ // 通过指针修改原变量
}

func main() {
	x := 10
	p := &x // p 的类型是 *int
	fmt.Println("x =", x, "地址 p =", p, "*p =", *p)

	*p = 20 // 解引用赋值
	fmt.Println("修改后 x =", x)

	addOne(&x)
	fmt.Println("addOne 后 x =", x)

	// new：为 int 分配堆内存，值为 0
	ptr := new(int)
	fmt.Println("*ptr 零值 =", *ptr)
	*ptr = 42
	fmt.Println("*ptr =", *ptr)

	// make vs new
	s := make([]int, 3)   // 切片，已可用
	m := make(map[int]int) // map，已可用
	fmt.Println("make slice =", s, "make map =", m)

	// 结构体指针：Go 自动解引用，p.Name 等价 (*p).Name
	type Point struct{ X, Y int }
	pt := &Point{X: 1, Y: 2}
	fmt.Println("pt.X =", pt.X)
}
