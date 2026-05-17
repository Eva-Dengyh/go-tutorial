// 本课目标：掌握 channel 的发送、接收、缓冲与 close。
//
// 知识点：
//   - ch := make(chan T) 无缓冲：发送与接收必须同时就绪
//   - make(chan T, n) 有缓冲：可存 n 个元素
//   - close(ch) 关闭；range ch 读尽剩余值
//   - 向已关闭 channel 发送会 panic
//
// 运行：go run ./chapter03_advanced/35_channel.go
package main

import "fmt"

func main() {
	// 无缓冲
	ch := make(chan int)
	go func() { ch <- 42 }()
	fmt.Println("无缓冲接收:", <-ch)

	// 有缓冲
	buf := make(chan string, 2)
	buf <- "a"
	buf <- "b"
	fmt.Println(<-buf, <-buf)

	// close + range
	nums := make(chan int, 3)
	nums <- 1
	nums <- 2
	nums <- 3
	close(nums)
	for n := range nums {
		fmt.Print(n, " ")
	}
	fmt.Println()
}
