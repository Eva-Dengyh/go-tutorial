// package calc 供单元测试示例使用。
package calc

// Add 两数相加
func Add(a, b int) int {
	return a + b
}

// Divide 除法，除数为 0 时返回 error
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
