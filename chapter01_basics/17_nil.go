// 本课目标：理解 nil 在不同类型上的含义与常见陷阱。
//
// 知识点：
//   - nil 是指针、切片、map、channel、func、interface 的零值
//   - nil 切片可 append；nil map 不能写入
//   - interface 为 nil 需「类型 + 值」都为 nil
//   - 与 null 不同：未初始化的引用类型多为 nil
//
// 运行：go run ./chapter01_basics/17_nil.go
package main

import "fmt"

func main() {
	var p *int
	var s []int
	var m map[string]int
	var ch chan int

	fmt.Println("指针 nil?", p == nil)
	fmt.Println("切片 nil?", s == nil, "len =", len(s))
	fmt.Println("map nil?", m == nil)
	fmt.Println("channel nil?", ch == nil)

	// nil 切片可以 append
	s = append(s, 1)
	fmt.Println("append 后 s =", s)

	// nil map 写入会 panic
	// m["k"] = 1
	m = make(map[string]int)
	m["k"] = 1
	fmt.Println("m =", m)

	// interface nil：持有一个具体类型的 nil 指针时，接口本身不是 nil
	var ptr *int
	var iface interface{} = ptr
	fmt.Println("iface == nil?", iface == nil) // false
	fmt.Println("ptr == nil?", ptr == nil)     // true
}
