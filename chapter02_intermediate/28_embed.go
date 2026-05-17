// 本课目标：掌握结构体嵌入（组合）与方法提升。
//
// 知识点：
//   - 嵌入字段：类型名作为字段名，无变量名
//   - 外层可直接访问内层导出字段和方法（提升）
//   - 组合优于继承：通过嵌入复用行为
//
// 运行：go run ./chapter02_intermediate/28_embed.go
package main

import "fmt"

type Engine struct {
	Power int
}

func (e Engine) Start() string {
	return fmt.Sprintf("引擎启动 %d 马力", e.Power)
}

type Car struct {
	Engine // 嵌入 Engine
	Brand  string
}

func main() {
	c := Car{
		Engine: Engine{Power: 200},
		Brand:  "示例牌",
	}
	// 方法提升：可直接 c.Start()
	fmt.Println(c.Brand, c.Start())
	fmt.Println("Power =", c.Power) // 字段提升
}
