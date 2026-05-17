package calc

import (
	"errors"
	"testing"
)

// 表驱动测试：多组输入输出写在表格里，便于扩展
func TestAdd(t *testing.T) {
	cases := []struct {
		name     string
		a, b, want int
	}{
		{"正数", 1, 2, 3},
		{"负数", -1, -2, -3},
		{"零", 0, 0, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Add(c.a, c.b)
			if got != c.want {
				t.Errorf("Add(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil || got != 5 {
		t.Fatalf("Divide(10,2) = %d, %v", got, err)
	}
	_, err = Divide(1, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Errorf("期望 ErrDivideByZero, got %v", err)
	}
}

// Benchmark 性能基准测试
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i+1)
	}
}
