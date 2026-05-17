// 本课目标：掌握 JSON 编解码与文件读写。
//
// 知识点：
//   - encoding/json：Marshal / Unmarshal
//   - 结构体字段标签 `json:"name"`
//   - os.ReadFile / os.WriteFile、os.Create
//
// 运行：go run ./chapter02_intermediate/33_json_fileio.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age,omitempty"` // 零值时省略
}

const demoFile = "user_demo.json"

func main() {
	u := User{Name: "小明", Email: "ming@example.com", Age: 20}

	// 编码为 JSON
	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON:\n", string(data))

	// 写入文件
	if err := os.WriteFile(demoFile, data, 0644); err != nil {
		panic(err)
	}
	defer os.Remove(demoFile)

	// 从文件读取并解码
	raw, err := os.ReadFile(demoFile)
	if err != nil {
		panic(err)
	}
	var u2 User
	if err := json.Unmarshal(raw, &u2); err != nil {
		panic(err)
	}
	fmt.Printf("从文件解析: %+v\n", u2)
}
