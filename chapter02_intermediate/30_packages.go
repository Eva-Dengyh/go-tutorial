// 本课目标：学会导入项目内的自定义子包。
//
// 知识点：
//   - 模块路径 + 目录路径 = import 路径
//   - 目录名不必与包名相同，但包名统一为一个
//   - 只导入副作用：import _ "包"
//
// 运行：go run ./chapter02_intermediate/30_packages.go
package main

import (
	"fmt"

	"github.com/Eva-Dengyh/go-tutorial/chapter02_intermediate/greet"
)

func main() {
	msg := greet.Hello("Go 学习者")
	fmt.Println(msg)
}
