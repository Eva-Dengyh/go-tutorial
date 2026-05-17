// 本课目标：掌握 context 的取消与超时。
//
// 知识点：
//   - context.Background() 根上下文
//   - WithCancel / WithTimeout / WithDeadline 派生子 context
//   - 监听 ctx.Done() 响应取消
//   - 不要用 context 传业务参数，只传请求域元数据
//
// 运行：go run ./chapter03_advanced/40_context.go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "退出:", ctx.Err())
			return
		default:
			fmt.Println(name, "工作中...")
			time.Sleep(80 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	go worker(ctx, "任务A")
	<-ctx.Done()
	time.Sleep(50 * time.Millisecond)
}
