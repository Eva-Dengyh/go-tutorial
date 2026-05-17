// 本课目标：掌握结构体定义、字段访问与构造模式。
//
// 知识点：
//   - type Name struct { ... } 定义结构体
//   - 字面量：按顺序或按字段名赋值
//   - 匿名嵌套：嵌入字段的方法可被提升
//   - Go 无 class，用结构体 + 方法表达数据
//
// 运行：go run ./chapter01_basics/14_struct.go
package main

import "fmt"

// Person 表示一个人的基本信息
type Person struct {
	Name string
	Age  int
}

// NewPerson 构造函数模式：返回初始化好的结构体指针
func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func main() {
	// 1. 按字段名初始化（推荐，可读性好）
	p1 := Person{Name: "小明", Age: 20}
	fmt.Println("p1 =", p1)

	// 2. 构造函数
	p2 := NewPerson("小红", 22)
	fmt.Println("p2 =", p2.Name, p2.Age)

	// 3. 指针与值：方法可修改字段时用指针
	p2.Age = 23
	fmt.Println("修改后 Age =", p2.Age)

	// 4. 匿名嵌套
	type Address struct {
		City string
	}
	type Employee struct {
		Person        // 匿名字段，字段可提升访问
		Address
		Dept string
	}
	emp := Employee{
		Person:  Person{Name: "张三", Age: 30},
		Address: Address{City: "上海"},
		Dept:    "研发",
	}
	// 提升：可直接 emp.Name 访问嵌入的 Person.Name
	fmt.Printf("员工 %s，%d 岁，%s，部门 %s\n",
		emp.Name, emp.Age, emp.City, emp.Dept)
}
