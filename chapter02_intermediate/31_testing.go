// 本课目标：了解如何运行本仓库中的单元测试。
//
// 知识点：
//   - 测试文件以 _test.go 结尾，包名一般为 被测包 或 被测包_test
//   - go test ./路径/ -v 运行测试
//   - 表驱动测试、t.Run 子测试、Benchmark 见 testing/calc_test.go
//
// 运行：go run ./chapter02_intermediate/31_testing.go
// 测试：go test ./chapter02_intermediate/testing/ -v
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("本课演示如何运行单元测试，请在本项目根目录执行：")
	fmt.Println("  go test ./chapter02_intermediate/testing/ -v")
	fmt.Println("  go test ./chapter02_intermediate/testing/ -bench=.")

	// 可选：自动跑一次测试（需在 go-tutorial 根目录）
	if _, err := os.Stat("chapter02_intermediate/testing/calc_test.go"); err == nil {
		cmd := exec.Command("go", "test", "./chapter02_intermediate/testing/", "-v")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("（若失败请确认在仓库根目录运行）")
		}
	}
}
