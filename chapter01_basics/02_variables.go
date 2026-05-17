// 本课目标：掌握 Go 中变量的多种声明方式，理解「零值」概念。
//
// 知识点：
//   - var 显式声明：可只写类型、只写初值、或两者都写
//   - := 短变量声明：函数内常用，左侧至少有一个新变量
//   - 零值：未显式赋值时，各类型有默认值（0、""、false、nil）
//   - 多变量：var a, b int 或 a, b := 1, 2
//
// 运行：go run ./chapter01_basics/02_variables.go
package main

import "fmt"

// 包级变量可用 var 声明；短声明 := 只能用在函数内
var globalCount int = 100

func main() {
	// 1. 完整声明：var 名 类型 = 值
	var age int = 25
	fmt.Println("age =", age)

	// 2. 类型推断：省略类型，由初值推断
	var name = "Go"
	fmt.Println("name =", name)

	// 3. 仅声明类型，使用零值（int 零值为 0）
	var score int
	fmt.Println("score 零值 =", score)

	// 4. 短变量声明（最常用）
	city := "北京"
	fmt.Println("city =", city)

	// 5. 多变量同时声明
	var x, y int = 1, 2
	a, b := 3, 4
	fmt.Printf("x=%d y=%d a=%d b=%d\n", x, y, a, b)

	// 6. 短声明至少有一个新变量名，否则编译错误
	// city := "上海"  // 错误：no new variables
	city, country := "上海", "中国" // 正确：country 是新变量
	fmt.Println(city, country)

	// 7. 常见类型的零值
	var (
		i    int
		f    float64
		s    string
		ok   bool
		ptr  *int
	)
	fmt.Printf("零值: int=%d float=%v string=%q bool=%v ptr=%v\n",
		i, f, s, ok, ptr)

	fmt.Println("包级变量 globalCount =", globalCount)
}
