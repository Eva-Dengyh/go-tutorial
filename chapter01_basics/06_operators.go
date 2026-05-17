// 本课目标：熟悉 Go 各类运算符及优先级。
//
// 知识点：
//   - 算术：+ - * / %（取余）
//   - 比较：== != < > <= >=
//   - 逻辑：&& || !（短路求值）
//   - 位运算：& | ^ << >> &^（按位清除）
//   - 赋值：= += -= 等
//
// 运行：go run ./chapter01_basics/06_operators.go
package main

import "fmt"

func main() {
	a, b := 10, 3
	fmt.Println("加减乘除余:", a+b, a-b, a*b, a/b, a%b)

	// 整数除法截断小数；需要浮点结果先转换
	fmt.Println("浮点除法:", float64(a)/float64(b))

	fmt.Println("比较:", a > b, a == b)

	x, y := true, false
	fmt.Println("逻辑与:", x && y)
	fmt.Println("逻辑或:", x || y)
	fmt.Println("逻辑非:", !x)

	// 位运算示例
	n := uint8(0b1010) // 10
	fmt.Printf("n=%08b\n", n)
	fmt.Printf("n<<1=%08b\n", n<<1)
	fmt.Printf("n&^1=%08b (清除最低位)\n", n&^1)

	// 短路：&& 左侧 false 不执行右侧；|| 左侧 true 不执行右侧
	result := false && sideEffect()
	fmt.Println("短路结果:", result, "sideEffect 未调用")

	_ = result
}

func sideEffect() bool {
	fmt.Println("sideEffect 被调用了")
	return true
}
