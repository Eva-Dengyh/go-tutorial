// 本课目标：用 Mutex 保护共享数据，避免竞态。
//
// 知识点：
//   - sync.Mutex：互斥锁，Lock/Unlock
//   - sync.RWMutex：读多写少时，RLock 允许多个读者
//   - go test -race 检测竞态
//
// 运行：go run ./chapter03_advanced/38_mutex.go
package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu sync.Mutex
	n  int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func main() {
	var c SafeCounter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println("计数 =", c.Value())
}
