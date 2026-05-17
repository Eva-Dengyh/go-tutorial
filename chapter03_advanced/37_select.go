// 本课目标：掌握 select 多路复用与超时。
//
// 知识点：
//   - select 监听多个 channel，哪个先就绪执行哪个 case
//   - default：非阻塞尝试
//   - 配合 time.After 实现超时
//
// 运行：go run ./chapter03_advanced/37_select.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(30 * time.Millisecond)
		ch1 <- "来自 ch1"
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "来自 ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("超时")
	}

	// default 非阻塞
	ch := make(chan int)
	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		fmt.Println("channel 暂无数据")
	}
}
