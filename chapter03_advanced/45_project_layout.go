// 本课目标：了解 Go 标准项目目录布局（非运行型，以注释说明为主）。
//
// 推荐结构：
//
//	myapp/
//	├── go.mod
//	├── cmd/
//	│   └── myapp/
//	│       └── main.go      # 可执行入口，可多个 cmd/xxx
//	├── internal/            # 私有代码，外部模块不可 import
//	│   └── service/
//	├── pkg/                 # 可被外部引用的库代码（慎用，确实需要再暴露）
//	├── api/                 # OpenAPI、protobuf 等
//	├── configs/
//	└── test/ 或 与源码同目录的 *_test.go
//
// 原则：
//   - 业务入口放 cmd/
//   - 不希望被引用的实现放 internal/
//   - 本教程用 chapter01_basics / chapter02_intermediate / chapter03_advanced 分难度
//
// 运行：go run ./chapter03_advanced/45_project_layout.go
package main

import "fmt"

func main() {
	fmt.Println("请阅读本文件顶部注释了解标准项目布局。")
	fmt.Println("本仓库教程目录：chapter01_basics → chapter02_intermediate → chapter03_advanced")
}
