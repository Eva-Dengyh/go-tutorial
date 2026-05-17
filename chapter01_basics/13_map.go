// 本课目标：掌握 map 的增删改查与 ok 模式。
//
// 知识点：
//   - map[K]V 是引用类型，必须用 make 或字面量初始化
//   - 取值 v, ok := m[key]：ok 表示 key 是否存在
//   - delete(m, key) 删除；遍历顺序随机
//   - 零值为 nil，不能对 nil map 写入
//
// 运行：go run ./chapter01_basics/13_map.go
package main

import "fmt"

func main() {
	// 1. 字面量创建
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("ages =", ages)

	// 2. make 创建
	scores := make(map[string]int)
	scores["语文"] = 90
	scores["数学"] = 95
	fmt.Println("scores =", scores)

	// 3. ok 模式：区分「不存在」与「值为零」
	if v, ok := scores["英语"]; ok {
		fmt.Println("英语", v)
	} else {
		fmt.Println("没有英语成绩")
	}

	// 4. 修改与删除
	scores["语文"] = 92
	delete(scores, "数学")
	fmt.Println("修改删除后 =", scores)

	// 5. 遍历（顺序不保证）
	fmt.Println("遍历:")
	for k, v := range scores {
		fmt.Printf("  %s: %d\n", k, v)
	}

	// 6. nil map 不能写入
	var m map[string]int
	fmt.Println("nil map == nil?", m == nil)
	// m["a"] = 1 // 运行时会 panic: assignment to entry in nil map
	m = make(map[string]int)
	m["a"] = 1
	fmt.Println("make 后可写入 m =", m)
}
