// 本课目标：掌握方法的定义，以及值接收者与指针接收者的区别。
//
// 知识点：
//   - 方法：带接收者的函数 func (接收者) 方法名(...)
//   - 值接收者：操作副本，不修改原值
//   - 指针接收者：可修改原值；大结构体避免复制
//   - 方法集：T 的方法含值接收者；*T 还含指针接收者
//
// 运行：go run ./chapter02_intermediate/21_methods.go
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// Area 值接收者：只读计算
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Scale 指针接收者：修改自身
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println("面积 =", rect.Area())

	rect.Scale(2)
	fmt.Printf("放大后 %.0f x %.0f, 面积 = %.0f\n",
		rect.Width, rect.Height, rect.Area())

	// 值类型也可调用指针接收者方法，Go 自动取地址
	r2 := Rectangle{1, 2}
	r2.Scale(3)
	fmt.Println("r2 =", r2)
}
