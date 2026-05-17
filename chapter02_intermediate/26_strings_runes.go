// 本课目标：理解 string、[]byte、[]rune 与 UTF-8。
//
// 知识点：
//   - string 是不可变 UTF-8 字节序列
//   - len(s) 是字节数，不是字符数
//   - range 遍历 string 得到 rune（码点）
//   - []byte(s) / string(b) 与 []rune(s) 转换
//
// 运行：go run ./chapter02_intermediate/26_strings_runes.go
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Go语言"
	fmt.Println("字节数 len =", len(s))
	fmt.Println("字符数 rune =", utf8.RuneCountInString(s))

	fmt.Print("按 rune 遍历: ")
	for i, r := range s {
		fmt.Printf("[%d]=%c ", i, r)
	}
	fmt.Println()

	// strings 包常用函数
	fmt.Println("Contains:", strings.Contains(s, "语言"))
	fmt.Println("ToUpper:", strings.ToUpper("hello"))

	// 与 []byte / []rune 互转
	b := []byte(s)
	runes := []rune(s)
	fmt.Println("[]byte 长度", len(b), "[]rune 长度", len(runes))
	fmt.Println("string([]rune):", string(runes))
}
