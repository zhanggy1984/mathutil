package mathutil

import "fmt"

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers.
func Subtract(a, b int) int {
	return a - b
}

// Multiply 返回两个整数的积（不兼容的 API 更改）
func Multiply(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("multiplication by zero")
	}
	return a * b, nil
}

// Divide returns the quotient of two integers.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
