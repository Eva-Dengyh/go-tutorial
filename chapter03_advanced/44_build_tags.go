// 本课目标：理解 build tags 条件编译。
//
// 知识点：
//   - 文件头 //go:build tag 指定编译条件
//   - go run -tags debug 启用带 debug tag 的文件
//   - 见 chapter03_advanced/buildtags/ 下 normal.go 与 debug.go
//
// 运行：
//
//	go run ./chapter03_advanced/44_build_tags.go
//	go run -tags debug ./chapter03_advanced/44_build_tags.go
package main

import (
	"fmt"

	"github.com/Eva-Dengyh/go-tutorial/chapter03_advanced/buildtags"
)

func main() {
	fmt.Println("当前编译模式:", buildtags.Mode())
}
