// 本课目标：掌握 WaitGroup、Once 的用法。
//
// 知识点：
//   - WaitGroup：Add、Done、Wait 等待一组 goroutine
//   - Once：保证函数只执行一次（如单例初始化）
//   - Cond：条件变量，复杂协调时使用
//
// 运行：go run ./chapter03_advanced/39_waitgroup.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("worker", id)
		}(i)
	}
	wg.Wait()

	var once sync.Once
	init := func() { fmt.Println("只初始化一次") }
	once.Do(init)
	once.Do(init)
}
