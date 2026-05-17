// 本课目标：掌握 range 遍历 slice、map、string、channel。
//
// 知识点：
//   - for i, v := range x 返回索引与值（map 为 key, value）
//   - 只要索引：for i := range s；只要值：for _, v := range s
//   - string 的 range 按 rune（Unicode 码点）遍历
//   - channel 关闭后 range 会读完剩余数据后退出
//
// 运行：go run ./chapter01_basics/16_range.go
package main

import "fmt"

func main() {
	// 1. 遍历 slice
	nums := []int{10, 20, 30}
	fmt.Println("slice:")
	for i, v := range nums {
		fmt.Printf("  [%d]=%d\n", i, v)
	}

	// 2. 遍历 map
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println("map:")
	for k, v := range m {
		fmt.Printf("  %s=%d\n", k, v)
	}

	// 3. 遍历 string（按 rune）
	s := "Go语言"
	fmt.Println("string rune:")
	for i, r := range s {
		fmt.Printf("  %d: %c (%U)\n", i, r, r)
	}

	// 4. 遍历 channel
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 关闭后才能被 range 读完
	fmt.Println("channel:")
	for v := range ch {
		fmt.Println(" ", v)
	}
}
