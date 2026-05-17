// 本课目标：理解接口的隐式实现与组合。
//
// 知识点：
//   - 接口定义方法集合，类型无需显式 implements
//   - 只要实现了接口全部方法，即满足该接口
//   - any（空接口 interface{}）可持有任意类型
//   - 接口可嵌入其他接口形成组合
//
// 运行：go run ./chapter02_intermediate/22_interfaces.go
package main

import "fmt"

// Speaker 会说话的接口
type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }

func (d Dog) Speak() string { return d.Name + ": 汪汪" }

type Cat struct{ Name string }

func (c Cat) Speak() string { return c.Name + ": 喵喵" }

func greet(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	greet(Dog{"小黑"})
	greet(Cat{"小花"})

	// 空接口 any 可装任意值
	var x any = 42
	fmt.Printf("any 存 int: %v, 类型 %T\n", x, x)
	x = "hello"
	fmt.Printf("any 存 string: %v\n", x)

	// 类型断言见 23_type_assertion.go
}
