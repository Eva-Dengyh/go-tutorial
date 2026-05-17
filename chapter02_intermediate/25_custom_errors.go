// 本课目标：自定义错误类型与错误包装。
//
// 知识点：
//   - 实现 Error() string 即可作为 error 使用
//   - 可在错误中携带额外字段（如 HTTP 状态码）
//   - errors.Unwrap / %w 形成错误链
//
// 运行：go run ./chapter02_intermediate/25_custom_errors.go
package main

import (
	"errors"
	"fmt"
)

// AppError 业务错误，带错误码
type AppError struct {
	Code    int
	Message string
	Cause   error
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error { return e.Cause }

func login(password string) error {
	if password == "" {
		return &AppError{Code: 400, Message: "密码不能为空"}
	}
	if password != "secret" {
		inner := errors.New("校验失败")
		return &AppError{Code: 401, Message: "登录失败", Cause: inner}
	}
	return nil
}

func main() {
	if err := login(""); err != nil {
		fmt.Println(err)
	}
	if err := login("wrong"); err != nil {
		fmt.Println(err)
		var appErr *AppError
		if errors.As(err, &appErr) {
			fmt.Println("错误码:", appErr.Code)
		}
	}
	if err := login("secret"); err == nil {
		fmt.Println("登录成功")
	}
}
