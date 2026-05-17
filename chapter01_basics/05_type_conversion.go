// 本课目标：掌握 Go 的类型转换与字符串/数字互转。
//
// 知识点：
//   - Go 没有隐式数值转换，必须 T(v) 显式转换
//   - strconv 包：Atoi、Itoa、ParseFloat、FormatInt 等
//   - fmt.Sprintf / fmt.Sscanf 做格式化与解析
//
// 运行：go run ./chapter01_basics/05_type_conversion.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 显式转换：不同类型即使底层相同也需转换
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

	// int ↔ string：不能用 string(i)，那是把 i 当码点转字符
	s := strconv.Itoa(123)
	fmt.Println("Itoa:", s)

	n, err := strconv.Atoi("456")
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	fmt.Println("Atoi:", n)

	// 浮点解析
	pi, err := strconv.ParseFloat("3.14", 64)
	if err == nil {
		fmt.Println("ParseFloat:", pi)
	}

	// fmt 格式化
	msg := fmt.Sprintf("用户%d余额%.2f", 1, 99.5)
	fmt.Println(msg)

	// 字符串转 []byte / []rune
	text := "Go语言"
	bytes := []byte(text)
	runes := []rune(text)
	fmt.Printf("字节数=%d 字符数=%d\n", len(bytes), len(runes))
}
