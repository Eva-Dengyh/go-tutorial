// 本课目标：了解 reflect 的基本用法与注意点。
//
// 知识点：
//   - reflect.TypeOf / reflect.ValueOf 获取类型与值
//   - 可修改可寻址的值：Elem、SetInt 等
//   - 反射性能较差，能静态类型解决就不用反射
//
// 运行：go run ./chapter03_advanced/41_reflection.go
package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Title string `json:"title"`
	Pages int    `json:"pages"`
}

func main() {
	b := Book{Title: "Go 入门", Pages: 300}
	t := reflect.TypeOf(b)
	v := reflect.ValueOf(b)
	fmt.Println("类型:", t.Name(), "字段数:", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)
		fmt.Printf("  %s (%s) = %v, tag=%s\n",
			field.Name, field.Type, val.Interface(), field.Tag.Get("json"))
	}

	// 修改需传指针且可寻址
	vp := reflect.ValueOf(&b).Elem()
	vp.FieldByName("Pages").SetInt(350)
	fmt.Println("修改后:", b)
}
