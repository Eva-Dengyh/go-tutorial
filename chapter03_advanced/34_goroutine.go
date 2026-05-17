// 本课目标：理解 goroutine 与并发、并行的区别。
//
// 知识点：
//   - go func() 启动 goroutine，与 main 并发执行
//   - 主 goroutine 结束则整个程序退出，子 goroutine 可能被强制结束
//   - 需要 sync.WaitGroup 或 channel 等待子任务（见后续课程）
//
// 运行：go run ./chapter03_advanced/34_goroutine.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func say(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go say("A", &wg)
	go say("B", &wg)
	wg.Wait()
	fmt.Println("全部完成")
}
