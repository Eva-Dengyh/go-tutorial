// 本课目标：掌握切片 slice 的创建、append、copy 与底层共享。
//
// 知识点：
//   - 切片是对数组的引用，包含 ptr、len、cap
//   - make([]T, len, cap) 创建切片；字面量 []T{...}
//   - append 可能触发扩容并分配新底层数组
//   - 切片赋值共享底层，修改会互相影响（未扩容时）
//
// 运行：go run ./chapter01_basics/12_slice.go
package main

import "fmt"

func main() {
	// 1. 字面量创建
	s1 := []int{1, 2, 3}
	fmt.Println("s1 =", s1, "len =", len(s1), "cap =", cap(s1))

	// 2. make：len 个零值，cap 为容量上限
	s2 := make([]int, 3, 5)
	fmt.Println("s2 =", s2, "cap =", cap(s2))

	// 3. 从数组截取
	arr := [5]int{10, 20, 30, 40, 50}
	s3 := arr[1:4] // 含下标 1,2,3
	fmt.Println("s3 =", s3)

	// 4. append 追加元素
	s1 = append(s1, 4, 5)
	fmt.Println("append 后 s1 =", s1)

	// 5. 共享底层：s3 与 arr 关联
	s3[0] = 999
	fmt.Println("修改 s3 后 arr =", arr)

	// 6. copy 复制元素（按较短长度）
	dst := make([]int, 2)
	n := copy(dst, s1)
	fmt.Println("copy 了", n, "个元素到 dst =", dst)

	// 7. nil 切片：长度为 0，可 append
	var s4 []int
	fmt.Println("nil 切片 len =", len(s4), "== nil?", s4 == nil)
	s4 = append(s4, 1)
	fmt.Println("append 后 s4 =", s4)
}
