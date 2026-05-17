// 本课目标：理解数组的声明、长度固定、值类型语义。
//
// 知识点：
//   - 数组长度是类型的一部分：[3]int 与 [5]int 是不同类型
//   - 值类型：赋值、传参会复制整个数组
//   - len 获取元素个数；索引从 0 开始
//   - 数组不能动态扩容，实际开发多用切片 slice
//
// 运行：go run ./chapter01_basics/11_array.go
package main

import "fmt"

func main() {
	// 1. 声明并指定初值
	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println("nums =", nums)

	// 2. 让编译器推断长度
	names := [...]string{"Go", "Rust", "Python"}
	fmt.Println("names 长度 =", len(names), "内容 =", names)

	// 3. 按索引访问与修改
	nums[0] = 100
	fmt.Println("修改后 nums[0] =", nums[0])

	// 4. 遍历
	fmt.Print("遍历: ")
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()

	// 5. 值类型：赋值会复制
	a := [2]int{1, 2}
	b := a
	b[0] = 99
	fmt.Println("a =", a, "b =", b) // a 不变

	// 6. 多维数组
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("matrix[1][2] =", matrix[1][2])
}
