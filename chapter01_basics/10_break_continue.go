// 本课目标：理解 break、continue 与标签 label。
//
// 知识点：
//   - break：跳出当前 for/switch/select
//   - continue：跳过本次循环剩余语句，进入下一轮
//   - 标签：break Label / continue Label 可跳出外层循环（慎用）
//   - goto 存在但不推荐，本课不展开
//
// 运行：go run ./chapter01_basics/10_break_continue.go
package main

import "fmt"

func main() {
	// break：找到第一个偶数就停
	fmt.Print("第一个偶数: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			break
		}
	}

	// continue：只打印奇数
	fmt.Print("奇数: ")
	for i := 1; i <= 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 标签：跳出外层循环
Outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j == 6 {
				fmt.Println("break Outer at", i, j)
				break Outer
			}
		}
	}
}
