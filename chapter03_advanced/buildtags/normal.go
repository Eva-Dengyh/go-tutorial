//go:build !debug

package buildtags

// Mode 返回当前编译模式
func Mode() string {
	return "release"
}
