// 本课目标：用 net/http 写最小 HTTP 服务与客户端请求。
//
// 知识点：
//   - http.HandleFunc 注册路由；ListenAndServe 监听
//   - http.Get / http.Post 发起请求
//   - ResponseWriter 写响应；r.URL.Query() 读查询参数
//
// 运行：go run ./chapter03_advanced/42_http_server.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "世界"
	}
	fmt.Fprintf(w, "你好, %s!", name)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	// 用 httptest 在内存中测 server，无需真起端口
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := http.Get(server.URL + "/hello?name=Go")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("状态:", resp.Status, "响应:", string(body))
}
