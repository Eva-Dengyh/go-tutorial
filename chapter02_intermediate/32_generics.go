// 本课目标：掌握 Go 1.18+ 泛型基础（类型参数与约束）。
//
// 知识点：
//   - func Name[T constraint](...) 类型参数 T
//   - 约束：interface{ ... } 或内置 comparable、any
//   - 泛型结构体、类型推断
//
// 运行：go run ./chapter02_intermediate/32_generics.go
package main

import "fmt"

// Number 约束：支持 + 的数字类型
type Number interface {
	~int | ~int64 | ~float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Stack 泛型栈
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() T {
	n := len(s.items) - 1
	v := s.items[n]
	s.items = s.items[:n]
	return v
}

func main() {
	fmt.Println("Max(3,7) =", Max(3, 7))
	fmt.Println("Max(1.2,3.4) =", Max(1.2, 3.4))

	st := &Stack[string]{}
	st.Push("a")
	st.Push("b")
	fmt.Println("Pop =", st.Pop())
}
