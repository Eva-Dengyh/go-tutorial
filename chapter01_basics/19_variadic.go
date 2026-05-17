// 本课目标：掌握可变参数函数 ... 的声明与调用。
//
// 知识点：
//   - 形参 ...T 表示接收 0 个或多个 T 类型参数，函数内为 []T
//   - 调用时可传多个实参，或展开切片：fn(slice...)
//   - 可变参数必须是参数列表最后一个
//
// 运行：go run ./chapter01_basics/19_variadic.go
package main

import "fmt"

// sum 计算任意个整数的和
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// join 用分隔符连接多个字符串
func join(sep string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for _, p := range parts[1:] {
		result += sep + p
	}
	return result
}

func main() {
	fmt.Println("sum() =", sum())
	fmt.Println("sum(1,2,3) =", sum(1, 2, 3))

	// 展开切片：把 []int 当作多个参数传入
	nums := []int{4, 5, 6}
	fmt.Println("sum(nums...) =", sum(nums...))

	fmt.Println(join("-", "Go", "语言", "教程"))
}
