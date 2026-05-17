// 本课目标：掌握 Go 唯一的循环关键字 for 的三种形式。
//
// 知识点：
//   - 经典三段式：for init; cond; post { }
//   - 仅条件：for cond { } 相当于 while
//   - 无限循环：for { } 需 break 退出
//   - Go 没有 while、do-while 关键字
//
// 运行：go run ./chapter01_basics/09_for_loop.go
package main

import "fmt"

func main() {
	// 1. 经典 for
	fmt.Print("0-4: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2. 仅条件（while 风格）
	sum := 0
	n := 1
	for n <= 5 {
		sum += n
		n++
	}
	fmt.Println("1+..+5 =", sum)

	// 3. 遍历 slice（也可用 range，见 16 课）
	nums := []int{10, 20, 30}
	fmt.Print("slice: ")
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()

	// 4. 嵌套循环打印乘法表一角
	fmt.Println("乘法表:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d×%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}
