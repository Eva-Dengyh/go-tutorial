// package greet 供 chapter02_intermediate/30_packages.go 演示自定义子包导入。
// 导出规则：首字母大写的标识符才能被其他包访问。
package greet

// Hello 向 name 问好（导出函数）
func Hello(name string) string {
	return "你好, " + name + "!"
}

// hello 小写未导出，包外无法调用
func hello() string {
	return "internal"
}
