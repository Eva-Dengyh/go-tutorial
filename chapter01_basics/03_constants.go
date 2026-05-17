// 本课目标：掌握常量与 iota 枚举写法。
//
// 知识点：
//   - const：编译期确定，不可修改
//   - 无类型常量：参与运算时可自动提升精度
//   - iota：在 const 块中从 0 递增，常用于枚举、位标志
//
// 运行：go run ./chapter01_basics/03_constants.go
package main

import "fmt"

const Pi = 3.14159
const AppName = "go-tutorial"

// iota 在 const 组内每行 +1，从 0 开始
const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
)

// 位掩码：用 iota 左移生成 1, 2, 4, 8...
const (
	Read   = 1 << iota // 1
	Write              // 2
	Execute            // 4
)

func main() {
	fmt.Println("Pi =", Pi, "App =", AppName)
	fmt.Println("周一枚举值 =", Monday)

	// 无类型常量与类型转换
	const bitSize = 64
	var n int = bitSize << 1 // 128
	fmt.Println("n =", n)

	// 权限组合：用 | 合并，用 & 检测
	perm := Read | Write
	fmt.Printf("有读权限: %v\n", perm&Read != 0)
	fmt.Printf("有执行权限: %v\n", perm&Execute != 0)

	// const 块内可跳过 iota 行，但序号仍递增
	const (
		_  = iota
		KB = 1 << (10 * iota) // 1024
		MB                    // 1048576
	)
	fmt.Println("1 MB =", MB, "字节")
}
