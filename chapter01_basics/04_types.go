// 本课目标：熟悉 Go 基本类型，理解类型定义与类型别名区别。
//
// 知识点：
//   - 布尔、整型、浮点、复数、字符串、byte、rune
//   - byte 是 uint8 别名；rune 是 int32 别名，表示 Unicode 码点
//   - type NewType BaseType 定义新类型（与别名不同）
//   - type Alias = BaseType 类型别名（Go 1.9+）
//
// 运行：go run ./chapter01_basics/04_types.go
package main

import "fmt"

// 定义新类型：UserID 与 int 是不同类型，不能直接混用
type UserID int

// 类型别名：与 int 完全等价
type Score = int

func main() {
	// 整型按范围选择：int8/16/32/64、uint、int（与平台相关，通常 64 位）
	var count int64 = 1_000_000 // 数字字面量可用 _ 分隔，提高可读性
	fmt.Println("count =", count)

	// 浮点：float32 / float64（默认推荐 float64）
	var price float64 = 19.99
	fmt.Println("price =", price)

	// bool
	var ok bool = true
	fmt.Println("ok =", ok)

	// string 是不可变 UTF-8 字节序列
	greeting := "你好"
	fmt.Println("len(字节) =", len(greeting)) // 中文通常占 3 字节/字

	// byte：单字节；rune：字符（码点）
	var ch byte = 'A'
	var r rune = '中'
	fmt.Printf("byte=%c rune=%c\n", ch, r)

	// 新类型需显式转换
	var id UserID = 42
	var n int = int(id) // UserID 不能赋给 int 而不转换
	fmt.Println("UserID =", id, "as int =", n)

	var s Score = 100
	var s2 int = s // 别名可直接互赋
	fmt.Println("Score =", s2)
}
