// 本课目标：了解 pipeline 与 fan-out 基本模式。
//
// 知识点：
//   - pipeline：阶段之间用 channel 串联
//   - fan-out：多个 worker 消费同一任务源
//   - 注意关闭 channel 的时机，避免泄漏
//
// 运行：go run ./chapter03_advanced/36_channel_patterns.go
package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	for n := range square(gen(1, 2, 3, 4)) {
		fmt.Print(n, " ") // 1 4 9 16
	}
	fmt.Println()
}
