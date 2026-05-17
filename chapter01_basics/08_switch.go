// 本课目标：掌握 switch 的多种用法。
//
// 知识点：
//   - 表达式 switch：switch x { case 1: ... }
//   - 无 tag switch：switch { case x > 0: } 类似 if-else 链
//   - case 可多个值：case 1, 2, 3:
//   - fallthrough：继续执行下一 case（默认自动 break）
//   - type switch 在进阶课 23 讲解
//
// 运行：go run ./chapter01_basics/08_switch.go
package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("周一")
	case 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("无效")
	}

	// 无 tag：条件写在 case 里
	score := 75
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// switch 也可带初始化
	switch hour := 14; {
	case hour < 12:
		fmt.Println("上午")
	case hour < 18:
		fmt.Println("下午")
	default:
		fmt.Println("晚上")
	}

	// fallthrough 慎用
	n := 2
	switch n {
	case 2:
		fmt.Print("二 ")
		fallthrough // 继续执行 case 3 的语句
	case 3:
		fmt.Println("（也进了 case 3）")
	default:
	}
}
