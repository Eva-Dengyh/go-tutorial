// 本课目标：掌握 Go 的错误处理惯例。
//
// 知识点：
//   - error 是内置接口：Error() string
//   - 多返回值 (_, err) 表示可能失败
//   - errors.Is / errors.As 判断错误链
//   - fmt.Errorf("... %w", err) 包装错误
//
// 运行：go run ./chapter02_intermediate/24_errors.go
package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrNotFound = errors.New("未找到")

func findUser(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("无效 id %d: %w", id, ErrNotFound)
	}
	return "用户" + fmt.Sprint(id), nil
}

func main() {
	name, err := findUser(1)
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	fmt.Println("找到:", name)

	_, err = findUser(-1)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("确认为 ErrNotFound")
	}

	// 打开不存在的文件
	_, err = os.Open("不存在的文件.txt")
	if err != nil {
		fmt.Println("os.Open 错误:", err)
	}
}
